package telebot_info

type APIResponseItemInfoList struct {
	Code    int                          `json:"code"`
	Message string                       `json:"message"`
	Data    *APIResponseItemInfoListData `json:"data"`
}
type APIResponseItemInfoListData struct {
	Items []TeleBotInfo `json:"items"`
	Total int64         `json:"total"`
}

type APIResponseItemInfo struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *TeleBotInfo `json:"data"`
}
