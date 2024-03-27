package user_func

import (
	"encoding/json"
	"fmt"

	"github.com/AnnonaOrg/osenv"

	model "github.com/AnnonaOrg/annona_bot/model/user_info"
	"github.com/AnnonaOrg/annona_bot/utils"
)

func DoAdd(item *model.UserInfo) (retText string, err error) {
	return doAddAPI(item)
}

func doAddAPI(req *model.UserInfo) (retText string, err error) {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/user/item/add"

	retBody, err := utils.DoPostJsonToOpenAPI(apiDomain, apiPath, apiToken, req)
	if err != nil {
		return "", fmt.Errorf("服务请求出错: %v", err)
	}
	var apiResponseItemInfo model.APIResponseItemInfo

	err = json.Unmarshal(retBody, &apiResponseItemInfo)
	if err != nil {
		return "", err
	}
	if apiResponseItemInfo.Code != 0 {
		err = fmt.Errorf("err msg: %s", apiResponseItemInfo.Message)
		return "", err
	}

	return retText, nil
}
