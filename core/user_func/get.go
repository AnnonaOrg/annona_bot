package user_func

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AnnonaOrg/annona_bot/core/utils"
	model "github.com/AnnonaOrg/annona_bot/model/user_info"
	"github.com/AnnonaOrg/osenv"
)

func DoGet(item *model.UserInfo) (retText string, err error) {
	return doGetAPI(item)
}

func doGetAPI(req *model.UserInfo) (retText string, err error) {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/user/item/get"

	retBody, err := utils.DoPostJsonToOpenAPI(apiDomain, apiPath, apiToken, req)
	if err != nil {
		return "", fmt.Errorf("服务请求出错: %v", err)
	}
	var apiResponseItemInfo model.APIResponseItemInfo

	err = json.Unmarshal(retBody, &apiResponseItemInfo)
	if err != nil {
		return "", fmt.Errorf("解析信息: %s ,失败: %v", string(retBody), err)
	}
	if apiResponseItemInfo.Code != 0 {
		return "", fmt.Errorf("请求失败: %s", apiResponseItemInfo.Message)
	}

	if apiResponseItemInfo.Data != nil { //&& apiResponseItemInfo.Data.Exp > 0
		retText = "ID: " + apiResponseItemInfo.Data.AccoundPlatformId
		inviterCode := ""
		if len(apiResponseItemInfo.Data.InviterCode) > 0 {
			inviterCode = apiResponseItemInfo.Data.InviterCode
		} else {
			inviterCode = apiResponseItemInfo.Data.AccoundPlatformId
		}
		expTimeStr := time.Unix(apiResponseItemInfo.Data.Exp, 0).Format(time.DateTime)
		retText = retText + "\n" +
			fmt.Sprintf("📢有效期: %s", expTimeStr) + "\n" +
			fmt.Sprintf("签到日期: %s", apiResponseItemInfo.Data.LastSignDate) + "\n" +
			fmt.Sprintf("邀请链接: t.me/%s?start=%s", req.TelegramStartBotUsername, inviterCode)
	} else {
		retText = "欢迎光临!"
	}

	return retText, nil
}
