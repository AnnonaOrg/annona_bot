package wspush

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/AnnonaOrg/annona_bot/core/utils"

	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

var fifoMap *FIFOMap

func init() {
	fifoMap = NewFIFOMap()
}
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}

	_, channelStr, ok := strings.Cut(r.URL.Path, "/ws/push_v2/")
	if len(channelStr) <= 0 || !ok {
		log.Debugf("收到非法推送: %s", r.URL.Path)
		return
	}
	log.Debugf("r.URL.Path: %s", r.URL.Path)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		log.Errorf("io.ReadAll(r.Body): %v", err)
		return
	}

	if err := PushMsgData(body); err != nil {
		fmt.Fprintf(w, "err: %v", err)
		log.Errorf("PushMsgData(%s): %v", string(body), err)
		return
	}
	fmt.Fprintf(w, "ok")
}

// 推送FeedMsg信息
func PushMsgData(data []byte) error {
	var msg FeedRichMsgModel
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Errorf("数据解析(%s)失败: %v", string(data), err)
		return err
	}
	if len(msg.MsgID) > 0 {
		msgID := "msgID_" + msg.MsgID
		if _, ok := fifoMap.Get(msgID); ok {
			return fmt.Errorf("msgID去重(%s)", msg.MsgID)
		} else {
			fifoMap.Set(msgID, true)
			if c := fifoMap.Count(); c > 100 {
				fifoMap.RemoveOldest()
			}
		}
	}

	return buildMsgDataAndSend(msg, SendMessage)
}

func buildMsgDataAndSend(msg FeedRichMsgModel,
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

func SendMessage(botToken string, reciverId int64, m interface{}, parseMode tele.ParseMode, noButton bool, button interface{}) error {
	bot, err := tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: false,
	})
	// common.Must(err)
	if err != nil {
		return err
	}

	reciver := &tele.User{
		ID: reciverId,
	}
	if noButton {
		if _, err := bot.Send(reciver, m, parseMode); err != nil {
			log.Printf("Send(%s,%d,%#v,%v) Msg Error: %v", botToken, reciverId, m, parseMode, err)
			return errors.New("send message failed")
		}
	} else {
		if _, err := bot.Send(reciver, m, parseMode, button); err != nil {
			log.Printf("Send(%s,%d,%#v,%v) Msg Error: %v", botToken, reciverId, m, parseMode, err)
			return errors.New("send message failed")
		}

	}

	return nil
}

type FeedRichMsgModel struct {
	Msgtype      string                `json:"msgtype"  form:"msgtype"` //rich text image video
	MsgID        string                `json:"msgID"  form:"msgID"`
	MsgTime      string                `json:"msgTime"  form:"msgTime"`
	Text         FeedRichMsgTextModel  `json:"text"  form:"text"`
	Image        FeedRichMsgImageModel `json:"image"  form:"image"`
	Video        FeedRichMsgVideoModel `json:"video"  form:"video"`
	Link         string                `json:"link"  form:"link"`
	LinkIsPublic bool                  `json:"linkIsPublic"  form:"linkIsPublic"`

	BotInfo  FeedRichMsgBotInfoModel  `json:"botInfo"  form:"botInfo"`
	ChatInfo FeedRichMsgChatInfoModel `json:"chatInfo"  form:"chatInfo"`
	FormInfo FeedRichMsgFormInfoModel `json:"formInfo"  form:"formInfo"`
	NoButton bool                     `json:"noButton" form:"noButton"`
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

type FeedRichMsgFormInfoModel struct {
	FormChatID   string `json:"formChatID" form:"formChatID"`
	FormSenderID string `json:"formSenderID" form:"formSenderID"`

	FormKeyworld string `json:"formKeyworld" form:"formKeyworld"`
}

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

// 先进先出（FIFO）的 map 实现
type FIFOMap struct {
	mu    sync.Mutex
	keys  []string
	items map[string]interface{}
}

func NewFIFOMap() *FIFOMap {
	return &FIFOMap{
		items: make(map[string]interface{}),
	}
}

func (m *FIFOMap) Set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 将 key 加入切片尾部，表示最近访问
	m.keys = append(m.keys, key)
	m.items[key] = value
}

func (m *FIFOMap) Get(key string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	value, ok := m.items[key]
	return value, ok
}

func (m *FIFOMap) RemoveOldest() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.keys) == 0 {
		return
	}

	oldestKey := m.keys[0]
	// 移除切片头部，表示最老的访问
	m.keys = m.keys[1:]
	delete(m.items, oldestKey)
}

func (m *FIFOMap) Count() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	return len(m.keys)
}
