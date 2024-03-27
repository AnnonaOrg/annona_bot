package blockword_func

import (
	"encoding/json"
	"fmt"

	"github.com/AnnonaOrg/osenv"

	model "github.com/AnnonaOrg/annona_bot/model/blockword_info"
	"github.com/AnnonaOrg/annona_bot/utils"
)

func DoDel(item *model.BlockworldInfo) (retText string, err error) {
	return doDelAPI(item)
}

func doDelAPI(req *model.BlockworldInfo) (retText string, err error) {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/blockword/item/del"

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

	retText = "删除成功"
	return retText, nil
}
