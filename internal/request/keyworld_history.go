package request

type KeyworldHistoryInfoRequest struct {
	ChatId             int64  `json:"chat_id" form:"chat_id" gorm:"column:chat_id;"`
	SenderId           int64  `json:"sender_id" form:"sender_id" gorm:"column:sender_id;"`
	SenderUsername     string `json:"sender_username" form:"sender_username" gorm:"column:sender_username;"`
	MessageId          int64  `json:"message_id" form:"message_id" gorm:"column:message_id;"`
	MessageLink        string `json:"message_link" form:"message_link" gorm:"column:message_link;"`
	MessageContentText string `json:"message_content_text" form:"message_content_text" gorm:"column:message_content_text;"`

	KeyWorld string `json:"key_world" form:"key_world" gorm:"column:key_world;"`

	// 核验请求id
	ById string `json:"by_id" form:"by_id" gorm:"-"`

	Page   int    `json:"-" form:"page" gorm:"-"`
	Size   int    `json:"-" form:"size" gorm:"-"`
	Filter string `json:"-" form:"filter" gorm:"-"`
}
