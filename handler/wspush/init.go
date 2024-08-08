package wspush

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/AnnonaOrg/annona_bot/core/response"
	"github.com/AnnonaOrg/annona_bot/core/service"
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
		// if _, ok := fifoMap.Get(msgID); ok {
		// 	return fmt.Errorf("msgID去重(%s)", msg.MsgID)
		// } else {
		// 	fifoMap.Set(msgID, true)
		// 	if c := fifoMap.Count(); c > 100 {
		// 		fifoMap.RemoveOldest()
		// 	}
		// }

		if _, ok := service.FIFOMapGet(msgID); ok {
			return fmt.Errorf("msgID去重(%s)", msg.MsgID)
		} else {
			service.FIFOMapSet(msgID, true)
			if c := service.FIFOMapCount(); c > 100 {
				service.FIFOMapRemoveOldest()
			}
		}
	}

	// return buildMsgDataAndSend(msg, SendMessage)
	return buildMsgDataAndSend(msg, service.SendMessage)
}

func buildMsgDataAndSend(msg response.FeedRichMsgResponse,
	sendMessage func(botToken string, reciverId int64, m interface{}, parseMode tele.ParseMode, noButton bool, button interface{}) error,
) error {
	reciverId := msg.ChatInfo.ToChatID
	botToken := msg.BotInfo.BotToken
	noButton := msg.NoButton

	selector := &tele.ReplyMarkup{}
	if !noButton {
		if len(msg.FormInfo.FormChatID) > 0 {
			noButton = false
		} else {
			noButton = true
		}
		btnSender := selector.Data("屏蔽号", "/block_formsenderid", msg.FormInfo.FormSenderID)
		btnChat := selector.Data("屏蔽群", "/block_formchatid", msg.FormInfo.FormChatID)
		btnByID := selector.Data("记录", "/by_formsenderid", msg.FormInfo.FormSenderID)
		btnByKeyworld := selector.Data("关键词", "/by_formkeyworld", msg.FormInfo.FormKeyworld)
		btnChatLink := selector.URL("私聊", "tg://user?id="+msg.FormInfo.FormSenderID)
		btnLink := selector.URL("定位消息", msg.Link)
		selector.Inline(
			selector.Row(btnLink, btnByID, btnChatLink),
			selector.Row(btnSender, btnChat, btnByKeyworld),
		)
	}

	switch msg.Msgtype {
	case "text":
		m := msg.Text.Content
		return sendMessage(botToken, reciverId, m, tele.ModeDefault, noButton, selector)

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
				m := msg.Text.Content
				return sendMessage(botToken, reciverId, m, tele.ModeDefault, noButton, selector)
			}
		default:
			return nil
		}
	default:
		return errors.New("msg type is not support,")
	}
}
