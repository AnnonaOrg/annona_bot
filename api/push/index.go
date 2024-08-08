package wspush

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/AnnonaOrg/annona_bot/common"
	"github.com/AnnonaOrg/annona_bot/core/utils"
	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}
	log.Println("r.URL.Path", r.URL.Path)
	// checkToken := r.Header.Get("Apiclient")
	// log.Println("checkToken", checkToken)
	// params := r.URL.Query()
	// userId, _ := strconv.Atoi(params.Get("to"))
	// msgText := params.Get("m")
	_, channelStr, ok := strings.Cut(r.URL.Path, "/ws/push/")
	if len(channelStr) <= 0 || !ok {
		log.Println("收到非法推送:(r.URL.Path)", r.URL.Path)
		return
	}

	body, err := io.ReadAll(r.Body)
	common.Must(err)

	if err := PushMsgData(body); err != nil {
		fmt.Fprintf(w, "err")
		return
	}
	fmt.Fprintf(w, "ok")
}

// 推送FeedMsg信息
func PushMsgData(data []byte) error {
	// fmt.Printf("\n解析接收到的数据data：%s\n", data)
	var msg FeedRichMsgModel
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Println(err, "解析数据失败:", string(data))
		return err
	}

	return buildMsgDataAndSend(msg, SendMessage)
}

func buildMsgDataAndSend(msg FeedRichMsgModel,
	sendMessage func(botToken string, reciverId int64, m interface{}, parseMode tele.ParseMode) error,
) error {
	reciverId := msg.ChatInfo.ToChatID
	botToken := msg.BotInfo.BotToken

	switch msg.Msgtype {
	case "text":
		m := msg.Text.Content
		return sendMessage(botToken, reciverId, m, tele.ModeDefault)

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
		return sendMessage(botToken, reciverId, m, tele.ModeHTML)

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
		return sendMessage(botToken, reciverId, m, tele.ModeHTML)

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
				return sendMessage(botToken, reciverId, m, tele.ModeHTML)
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
				return sendMessage(botToken, reciverId, m, tele.ModeHTML)
			}
		case len(msg.Text.Content) > 0:
			{
				m := msg.Text.Content
				return sendMessage(botToken, reciverId, m, tele.ModeDefault)
			}
		default:
			return nil
		}
	default:
		return errors.New("msg type is not support,")
	}
}

func SendMessage(botToken string, reciverId int64, m interface{}, parseMode tele.ParseMode) error {
	bot, err := tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: false,
	})
	common.Must(err)

	reciver := &tele.User{
		ID: reciverId,
	}

	if _, err := bot.Send(reciver, m, parseMode); err != nil {
		log.Printf("Send(%s,%d,%#v,%v) Msg Error: %v", botToken, reciverId, m, parseMode, err)
		return errors.New("send message failed")
	}

	return nil
}

type FeedRichMsgModel struct {
	Msgtype string                `json:"msgtype"  form:"msgtype"` //rich text image video
	MsgID   string                `json:"msgID"  form:"msgID"`
	MsgTime string                `json:"msgTime"  form:"msgTime"`
	Text    FeedRichMsgTextModel  `json:"text"  form:"text"`
	Image   FeedRichMsgImageModel `json:"image"  form:"image"`
	Video   FeedRichMsgVideoModel `json:"video"  form:"video"`
	Link    string                `json:"link"  form:"link"`

	BotInfo  FeedRichMsgBotInfoModel  `json:"botInfo"  form:"botInfo"`
	ChatInfo FeedRichMsgChatInfoModel `json:"chatInfo"  form:"chatInfo"`
	// FormInfo FeedRichMsgFormInfoModel `json:"formInfo"  form:"formInfo"`
	// NoButton bool                     `json:"noButton" form:"noButton"`
}
type FeedRichMsgTextModel struct {
	Content         string `json:"content"  form:"content"`
	ContentEx       string `json:"contentEx"  form:"contentEx"`
	ContentExPic    string `json:"contentExPic"  form:"contentExPic"`
	ContentMarkdown string `json:"contentMarkdown"  form:"contentMarkdown"`
}
type FeedRichMsgImageModel struct {
	PicURL   string `json:"picURL"  form:"picURL"`
	FilePath string `json:"filePath"  form:"filePath"`
	// (Optional)
	Caption string `json:"caption,omitempty"`
}
type FeedRichMsgVideoModel struct {
	FileURL  string `json:"fileURL"  form:"fileURL"`
	FilePath string `json:"filePath"  form:"filePath"`
	// (Optional)
	Caption string `json:"caption,omitempty"`
}
type FeedRichMsgBotInfoModel struct {
	BotToken string `json:"botToken" form:"botToken"`
}
type FeedRichMsgChatInfoModel struct {
	ToChatID int64 `json:"toChatID" form:"toChatID"`
}

// type FeedRichMsgFormInfoModel struct {
// 	FormChatID string `json:"formChatID" form:"formChatID"`
// }

func (msg *FeedRichMsgModel) ToString() (res string) {
	res = fmt.Sprintf(
		"msgID:%s\n"+"msgType:%s\n"+"msgTime:%s\n"+"toChatID:%d",
		msg.MsgID, msg.Msgtype, msg.MsgTime, msg.ChatInfo.ToChatID,
	)
	if len(msg.Text.Content) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Text.Content)
	}
	if len(msg.Image.PicURL) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Image.PicURL)
	}
	if len(msg.Video.FileURL) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Video.FileURL)
	}

	return
}
