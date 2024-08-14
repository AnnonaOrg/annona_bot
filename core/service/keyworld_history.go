package service

import (
	"encoding/json"
	"fmt"

	"github.com/AnnonaOrg/annona_bot/core/constvar"
	"github.com/AnnonaOrg/annona_bot/core/log"
	"github.com/AnnonaOrg/annona_bot/core/request"
	"github.com/AnnonaOrg/annona_bot/core/response"
	"github.com/AnnonaOrg/annona_bot/core/utils"
	"github.com/AnnonaOrg/osenv"
)

func GetListKeyworldHistory(req *request.KeyworldHistoryInfoRequest) ([]response.KeyworldHistoryInfoItem, error) {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/keyword_history/list"
	log.Debugf("%s%s", apiDomain, apiPath)
	retBody, err := utils.DoPostJsonToOpenAPI(apiDomain, apiPath, apiToken, req)
	if err != nil {
		log.Errorf("DoPostJsonToOpenAPI(%s,%s,%s,%+v): %v", apiDomain, apiPath, apiToken, req, err)
		return nil, fmt.Errorf("服务请求出错: %v", err)
	}
	// log.Debugf("retBody: %s", retBody)
	var apiResponse response.KeyworldHistoryInfoResponse

	err = json.Unmarshal(retBody, &apiResponse)
	if err != nil {
		log.Errorf("Unmarshal(%s): %v", string(retBody), err)
		return nil, fmt.Errorf("解析信息: %s ,失败: %v", string(retBody), err)
	}
	if apiResponse.Code != 0 {
		log.Errorf("响应状态异常: %+v", apiResponse)
		return nil, fmt.Errorf("err msg: %s", apiResponse.Message)
	}
	var list []response.KeyworldHistoryInfoItem
	for _, v := range apiResponse.Data.Items {
		vc := v
		var item response.KeyworldHistoryInfoItem
		item.Note = vc.Note

		item.MessageContentText = vc.MessageContentText
		item.SenderUsername = vc.SenderUsername
		item.SenderId = vc.SenderId
		item.KeyWorld = vc.KeyWorld
		item.Total = vc.Total
		item.MessageLink = vc.MessageLink
		list = append(list, item)
	}
	return list, nil
}
func GetListKeyworldHistoryWithSenderID(senderID int64, page int) (string, error) {
	req := &request.KeyworldHistoryInfoRequest{}
	req.SenderId = senderID
	req.Page = page
	req.Size = 50
	retList, err := GetListKeyworldHistory(req)
	if err != nil {
		return constvar.ERR_MSG_Server, err
	}
	retText := ""
	senderUsername := ""
	for k, v := range retList {
		if len(v.SenderUsername) > 0 && len(senderUsername) == 0 {
			senderUsername = "@" + v.SenderUsername
		}
		messageContentText := utils.GetStringRuneN(v.MessageContentText, 20)
		retText = fmt.Sprintf("%s\n %d. %s", retText,
			k, messageContentText,
		)
	}
	if len(retText) > 0 {
		retText = fmt.Sprintf("#ID%d ", senderID) + senderUsername + retText
	}
	return retText, nil
}

func GetListKeyworldHistoryWithKeyworld(keyworld string, page int) (string, error) {
	req := &request.KeyworldHistoryInfoRequest{}
	req.KeyWorld = keyworld
	req.Page = page
	req.Size = 50
	retList, err := GetListKeyworldHistory(req)
	if err != nil {
		return constvar.ERR_MSG_Server, err
	}
	retText := ""
	for k, v := range retList {
		text := v.Note
		if len(text) == 0 || len(v.MessageLink) == 0 {
			log.Debugf("retList: %d. %+v", k, v)
			continue
		}

		// senderUsername := v.SenderUsername
		// if len(senderUsername) > 0 {
		// 	senderUsername = "@" + senderUsername
		// } else {
		// 	continue
		// }
		// retText = fmt.Sprintf("%s\n %d. %s", retText,
		// 	k, senderUsername,
		// )
		retTextTmp := fmt.Sprintf("%s\n %d. %s", retText,
			k, text,
		)
		if len(retTextTmp) > 4096-200 || len(retText) > 3000 {
			break
		}
		retText = retTextTmp
	}
	retText = "关键词 #" + keyworld + ": " + retText

	log.Debugf("GetListKeyworldHistoryWithKeyworld(%s): %s", keyworld, retText)
	return retText, nil
}
