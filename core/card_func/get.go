package card_func

import (
	"encoding/json"
	"fmt"

	"github.com/AnnonaOrg/annona_bot/core/utils"
	model "github.com/AnnonaOrg/annona_bot/model/card_info"
	"github.com/AnnonaOrg/osenv"
)

func DoGet(item *model.CardInfo) (retText string, err error) {
	return doGetAPI(item)
}

func doGetAPI(req *model.CardInfo) (retText string, err error) {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/card/item/get"

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

	if apiResponseItemInfo.Data != nil && apiResponseItemInfo.Data.Exp > 0 {
		retText = apiResponseItemInfo.Data.CardUUID
		expTimeStr := fmt.Sprintf("%d 天", apiResponseItemInfo.Data.Exp) //time.Unix(apiResponseItemInfo.Data.Exp+time.Now().Unix(), 0).Format(time.DateTime)

		switch {
		case apiResponseItemInfo.Data.Stat == 1:
			retText =
				retText + "\n" +
					fmt.Sprintf("有效期: %s", expTimeStr) + "\n" +
					"已登记"
		case apiResponseItemInfo.Data.Stat == 2:
			retText =
				retText + "\n" +
					fmt.Sprintf("有效期: %s", expTimeStr) + "\n" +
					"待使用"
		case apiResponseItemInfo.Data.Stat == 3:
			retText = retText + "\n" + "已核销"
		default:
			retText = retText + "\n" + "无效卡"
		}
	} else {
		retText = "信息错误!"
	}

	return retText, nil
}
