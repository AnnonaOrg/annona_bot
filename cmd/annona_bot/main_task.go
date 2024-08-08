package main

import (
	"fmt"
	"time"

	"github.com/AnnonaOrg/osenv"
	"github.com/AnnonaOrg/annona_bot/core/constvar"
	"github.com/AnnonaOrg/annona_bot/core/log"
	"github.com/AnnonaOrg/annona_bot/core/service"
	"github.com/AnnonaOrg/annona_bot/core/utils"
)

func mainTask() {
	if osenv.GetBotTelegramWebhookURL() == "" {
		go mainBot()
	} else {
		go service.SetBotFatherWebhook()
	}
}

// 自检openAPI服务是否正常运行
func pingServer() error {
	apiURL := osenv.GetServerUrl()
	for i := 0; i < 10; i++ {

		if utils.CheckPingBaseURL(apiURL) {
			return nil
		}

		log.Debugf(
			"(%s)等待自检, 1秒后重试(%d) %s",
			constvar.APPName(), i, apiURL,
		)
		time.Sleep(time.Second * 2)
	}
	return fmt.Errorf(
		"(%s)自检失败 %s.",
		constvar.APPName(), apiURL,
	)
}
