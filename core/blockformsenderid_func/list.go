package blockformsenderid_func

import (
	"encoding/json"
	"fmt"

	"github.com/AnnonaOrg/annona_bot/core/utils"
	model "github.com/AnnonaOrg/annona_bot/model/blockformsenderid_info"
	"github.com/AnnonaOrg/osenv"
)

func DoList(item *model.BlockformsenderidInfo) (retText string, err error) {
	return doListAPI(item)
}

func doListAPI(req *model.BlockformsenderidInfo) (retText string, err error) {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/blockformsenderid/list"

	retBody, err := utils.DoPostJsonToOpenAPI(apiDomain, apiPath, apiToken, req)
	if err != nil {
		return "", fmt.Errorf("服务请求出错: %v", err)
	}
	var apiResponseItemInfoList model.APIResponseItemInfoList

	err = json.Unmarshal(retBody, &apiResponseItemInfoList)
	if err != nil {
		return "", fmt.Errorf("解析信息: %s ,失败: %v", string(retBody), err)
	}
	if apiResponseItemInfoList.Code != 0 {
		return "", fmt.Errorf("请求失败: %s", apiResponseItemInfoList.Message)
	}
	if apiResponseItemInfoList.Data != nil && apiResponseItemInfoList.Data.Total > 0 {
		textTmp := fmt.Sprintf("共计 %d 个\n", apiResponseItemInfoList.Data.Total)
		for i, v := range apiResponseItemInfoList.Data.Items {
			vc := v
			textTmp = textTmp + " " + "\n" +
				fmt.Sprintf("%d: %s ", i+1, vc.KeyWorld)
		}
		retText = textTmp
	} else {
		retText = "还未添加？"
	}

	return retText, nil
}
