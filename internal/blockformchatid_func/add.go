package blockformchatid_func

import (
	"encoding/json"
	"fmt"

	"github.com/AnnonaOrg/osenv"

	model "github.com/AnnonaOrg/annona_bot/model/blockformchatid_info"
	"github.com/AnnonaOrg/annona_bot/utils"
)

func DoAdd(item *model.BlockformchatidInfo) (retText string, err error) {
	return doAddAPI(item)
}

func doAddAPI(req *model.BlockformchatidInfo) (retText string, err error) {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/blockformchatid/item/add"

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
		err = fmt.Errorf("err msg: %s", apiResponseItemInfo.Message)
		return "", err
	}
	retText = "恭喜，添加成功"
	return retText, nil
}
