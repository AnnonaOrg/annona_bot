package response

type KeyworldHistoryInfoResponse struct {
	// Code    int         `json:"status"` //返回 0，表示当前接口正确返回，否则按错误请求处理；
	// Message string      `json:"msg"`    //返回接口处理信息，主要用于表单提交或请求失败时的 toast 显示；
	// Data    interface{} `json:"data"`   //必须返回一个具有 key-value 结构的对象。
	Code    int                             `json:"code"`
	Message string                          `json:"message"`
	Data    KeyworldHistoryInfoListResponse `json:"data"`
}
type KeyworldHistoryInfoListResponse struct {
	Total int64                     `json:"total" form:"total"`
	Items []KeyworldHistoryInfoItem `json:"items" form:"items"`
	// AdList interface{} `json:"ad_list,omitempty" form:"ad_list"`
}

type KeyworldHistoryInfoItem struct {
	ChatId             int64  `json:"chat_id" form:"chat_id" gorm:"column:chat_id;"`
	SenderId           int64  `json:"sender_id" form:"sender_id" gorm:"column:sender_id;"`
	SenderUsername     string `json:"sender_username" form:"sender_username" gorm:"column:sender_username;"`
	MessageId          int64  `json:"message_id" form:"message_id" gorm:"column:message_id;"`
	MessageLink        string `json:"message_link" form:"message_link" gorm:"column:message_link;"`
	MessageContentText string `json:"message_content_text" form:"message_content_text" gorm:"column:message_content_text;"`

	KeyWorld string `json:"key_world" form:"key_world" gorm:"column:key_world;"`
}
