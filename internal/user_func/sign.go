package user_func

import (
	"encoding/json"
	"fmt"

	"github.com/AnnonaOrg/osenv"

	model "github.com/AnnonaOrg/annona_bot/model/user_info"
	"github.com/AnnonaOrg/annona_bot/utils"
)

// 签到
func DoSign(item *model.UserInfo) (retText string, err error) {
	return doSignAPI(item)
}

func doSignAPI(req *model.UserInfo) (retText string, err error) {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/user/item/sign" //+ req.LastCardUUID

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

	retText = "恭喜，签到成功"
	return retText, nil
}
