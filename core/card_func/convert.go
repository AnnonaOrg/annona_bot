package card_func

import (
	"encoding/json"
	"fmt"

	"github.com/AnnonaOrg/annona_bot/core/utils"
	model "github.com/AnnonaOrg/annona_bot/model/card_info"
	"github.com/AnnonaOrg/osenv"
)

func DoConvert(item *model.CardInfo) (retText string, err error) {
	return doConvertAPI(item)
}

func doConvertAPI(req *model.CardInfo) (retText string, err error) {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/card/item/convert"

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
		err = fmt.Errorf("err msg: %s", apiResponseItemInfoList.Message)
		return "", err
	}

	// expTimeStr := time.Unix(apiResponseItemInfo.Data.Exp, 0).Format(time.DateTime)
	if apiResponseItemInfoList.Data != nil && apiResponseItemInfoList.Data.Total > 0 {
		textTmp := ""
		expTimeStr := ""
		for _, v := range apiResponseItemInfoList.Data.Items {
			vc := v
			if expTimeStr == "" {
				// expTimeStr = time.Unix(vc.Exp, 0).Format(time.DateTime)
				expTimeStr = fmt.Sprintf("%d 天", vc.Exp)
			}
			textTmp = textTmp + "\n" +
				vc.CardUUID + "  " + "有效期: " + expTimeStr
		}

		if len(textTmp) > 0 {
			textTmp = textTmp + "\n" +
				fmt.Sprintf("本次共计生成: %d ", apiResponseItemInfoList.Data.Total)
			retText = textTmp
		}

	} else {
		return "", fmt.Errorf("状态正常，本次未生成任何可用卡片")
	}

	return retText, nil
}
