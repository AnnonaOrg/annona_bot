package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/AnnonaOrg/annona_bot/core/response"
	"github.com/AnnonaOrg/annona_bot/core/utils"
	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

// 推送FeedMsg信息
func PushMsgData(data []byte) error {
	var msg response.FeedRichMsgResponse // FeedRichMsgModel
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Errorf("数据解析(%s)失败: %v", string(data), err)
		return err
	}
	if len(msg.MsgID) > 0 {
		msgID := "msgID_" + msg.MsgID

		if _, ok := FIFOMapGet(msgID); ok {
			return fmt.Errorf("msgID去重(%s)", msg.MsgID)
		} else {
			FIFOMapSet(msgID, true)
			if c := FIFOMapCount(); c > 100 {
				FIFOMapRemoveOldest()
			}
		}
	}

	return buildMsgDataAndSend(msg, SendMessage)
}

func buildMsgDataAndSend(msg response.FeedRichMsgResponse,
	sendMessage func(botToken string, reciverId int64, m interface{}, parseMode tele.ParseMode, noButton bool, button interface{}) error,
) error {
	reciverId := msg.ChatInfo.ToChatID
	botToken := msg.BotInfo.BotToken
	noButton := msg.NoButton

	selector := &tele.ReplyMarkup{}
	selector2 := &tele.ReplyMarkup{}
	if !noButton {
		if len(msg.FormInfo.FormChatID) > 0 {
			noButton = false
		} else {
			noButton = true
		}

		btnSender := selector.Data("屏蔽号", "/block_formsenderid", msg.FormInfo.FormSenderID)
		btnChat := selector.Data("屏蔽群", "/block_formchatid", msg.FormInfo.FormChatID)
		btnByKeyworld := selector.Data("关键词", "/by_formkeyworld", msg.FormInfo.FormKeyworld)

		btnLink := selector.URL("定位消息", msg.Link)
		if len(msg.Link) == 0 {
			if len(msg.FormInfo.FormChatUsername) > 0 {
				btnLink = selector.URL("定位消息", "https://t.me/"+msg.FormInfo.FormChatUsername)
			}
		}
		btnByID := selector.Data("记录", "/by_formsenderid", msg.FormInfo.FormSenderID)
		btnChatLink := selector.URL("私聊", "tg://user?id="+msg.FormInfo.FormSenderID)
		if len(msg.FormInfo.FormSenderUsername) > 0 {
			btnChatLink = selector.URL("私聊", "https://t.me/"+msg.FormInfo.FormSenderUsername)
		}

		selector.Inline(
			selector.Row(btnSender, btnChat, btnByKeyworld),
			selector.Row(btnLink, btnByID, btnChatLink),
		)
		selector2.Inline(
			selector2.Row(btnSender, btnChat, btnByKeyworld),
			selector2.Row(btnLink, btnByID),
		)
		// log.Debugf("btnLink: %+v", btnLink)
		// log.Debugf("btnByID: %+v", btnByID)
		// log.Debugf("btnChatLink: %+v", btnChatLink)
	}

	messageContentText := msg.Text.Content
	if len(msg.Text.ContentEx) > 0 {
		messageContentText = msg.Text.ContentEx
	} else if len(msg.Text.ContentHtml) > 0 {
		// messageContentText = msg.Text.ContentHtml
		textTmp := ""
		if len(msg.FormInfo.FormSenderTitle) > 0 {
			textTmp = "发送人:" + msg.FormInfo.FormSenderTitle
			if len(msg.FormInfo.FormSenderUsername) > 0 {
				textTmp = textTmp + " @" + msg.FormInfo.FormSenderUsername
			}
		}

		textTmp2 := ""
		if len(msg.FormInfo.FormChatTitle) > 0 {

			textTmp2 = "来源:"
			if len(msg.FormInfo.FormChatUsername) > 0 {
				textTmp2 = textTmp2 +
					fmt.Sprintf("<a href=\"https://t.me/%s\">%s</a>",
						msg.FormInfo.FormChatUsername, msg.FormInfo.FormChatTitle,
					)
			} else if len(msg.Link) > 0 {
				textTmp2 = textTmp2 +
					fmt.Sprintf("<a href=\"%s\">%s</a>",
						msg.Link, msg.FormInfo.FormChatTitle,
					)
			} else {
				textTmp2 = textTmp2 + msg.FormInfo.FormChatTitle
			}
		}

		if len(textTmp) > 0 {
			messageContentText = messageContentText + "\n" + textTmp
		}
		if len(textTmp2) > 0 {
			messageContentText = messageContentText + "\n" + textTmp2
		}
		messageContentText = messageContentText + "\n" + "#ID" + msg.FormInfo.FormSenderID
	}
	// fmt.Println("messageContentText", messageContentText)
	log.Debugf("待发送消息:%s", messageContentText)

	switch msg.Msgtype {
	case "text":
		m := messageContentText
		return sendMessage(botToken, reciverId, m, tele.ModeHTML, noButton, selector)

	case "video":
		m := new(tele.Video)
		m.File = tele.FromURL(msg.Video.FileURL)
		if len(msg.Video.Caption) > 0 {
			m.Caption = msg.Video.Caption
		}
		if len(m.Caption) > 0 {
			if captionTmp, err := utils.UrlRegMatchReplaceToTGHTML(m.Caption); err != nil {
			} else {
				m.Caption = captionTmp
			}
		}
		return sendMessage(botToken, reciverId, m, tele.ModeHTML, noButton, selector)

	case "image":
		m := new(tele.Photo)
		m.File = tele.FromURL(msg.Image.PicURL)
		if len(msg.Image.Caption) > 0 {
			m.Caption = msg.Image.Caption
		}
		if len(m.Caption) > 0 {
			if captionTmp, err := utils.UrlRegMatchReplaceToTGHTML(m.Caption); err != nil {
			} else {
				m.Caption = captionTmp
			}
		}
		return sendMessage(botToken, reciverId, m, tele.ModeHTML, noButton, selector)

	case "rich":
		switch {
		case len(msg.Video.FileURL) > 0 && strings.HasPrefix(msg.Video.FileURL, "http"):
			{
				m := new(tele.Video)
				m.File = tele.FromURL(msg.Video.FileURL)
				if len(msg.Video.Caption) > 0 {
					m.Caption = msg.Video.Caption
				} else if len(msg.Text.Content) > 0 {
					m.Caption = msg.Text.Content
				}
				if len(m.Caption) > 0 {
					if captionTmp, err := utils.UrlRegMatchReplaceToTGHTML(m.Caption); err != nil {
					} else {
						m.Caption = captionTmp
					}
				}
				return sendMessage(botToken, reciverId, m, tele.ModeHTML, noButton, selector)
			}
		case len(msg.Image.PicURL) > 0 && strings.HasPrefix(msg.Image.PicURL, "http"):
			{
				m := new(tele.Photo)
				m.File = tele.FromURL(msg.Image.PicURL)
				if len(msg.Image.Caption) > 0 {
					m.Caption = msg.Image.Caption
				} else if len(msg.Text.Content) > 0 {
					m.Caption = msg.Text.Content
				}
				if len(m.Caption) > 0 {
					if captionTmp, err := utils.UrlRegMatchReplaceToTGHTML(m.Caption); err != nil {
					} else {
						m.Caption = captionTmp
					}
				}
				return sendMessage(botToken, reciverId, m, tele.ModeHTML, noButton, selector)
			}
		case len(msg.Text.Content) > 0:
			{
				m := messageContentText //msg.Text.Content

				err := sendMessage(botToken, reciverId, m, tele.ModeHTML, noButton, selector)
				if err != nil {
					if strings.Contains(err.Error(), "BUTTON_USER_INVALID") && IsRetryPushMsgEnable() {
						if len(msg.FormInfo.FormSenderUsername) == 0 && (len(msg.Link) > 0 || len(msg.FormInfo.FormChatUsername) > 0) {
							return sendMessage(botToken, reciverId, m, tele.ModeHTML, noButton, selector2)
						}
					}
					return err
				}
				return nil
			}
		default:
			return nil
		}
	default:
		return errors.New("msg type is not support,")
	}
}
