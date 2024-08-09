package response

import (
	"fmt"
)

type FeedRichMsgResponse struct {
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
	FormMessageID int64 `json:"formMessageID" form:"formMessageID"`

	FormChatID       string `json:"formChatID" form:"formChatID"`
	FormChatUsername string `json:"formChatUsername" form:"formChatUsername"`
	FormChatTitle    string `json:"formChatTitle" form:"formChatTitle"`

	FormSenderID       string `json:"formSenderID" form:"formSenderID"`
	FormSenderUsername string `json:"formSenderUsername" form:"formSenderUsername"`
	FormSenderTitle    string `json:"formSenderTitle" form:"formSenderTitle"`

	FormKeyworld string `json:"formKeyworld" form:"formKeyworld"`
}

func (msg *FeedRichMsgResponse) ToString() (res string) {
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
