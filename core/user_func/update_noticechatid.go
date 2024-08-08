package user_func

import (
	"encoding/json"
	"fmt"

	"github.com/AnnonaOrg/annona_bot/core/utils"
	model "github.com/AnnonaOrg/annona_bot/model/user_info"
	"github.com/AnnonaOrg/osenv"
)

func DoUpdateNoticeChatId(item *model.UserInfo) (retText string, err error) {
	return doUpdateNoticeChatIdAPI(item)
}

func doUpdateNoticeChatIdAPI(req *model.UserInfo) (retText string, err error) {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/user/item/updatenoticechatid/" + fmt.Sprintf("%d", req.TelegramNoticeChatId)

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

	retText = "恭喜，更新成功"
	return retText, nil
}
