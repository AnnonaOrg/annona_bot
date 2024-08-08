package main

import (
	"time"

	_ "github.com/AnnonaOrg/annona_bot/cmd/annona_bot/distro/all"
	"github.com/AnnonaOrg/annona_bot/common"
	_ "github.com/AnnonaOrg/annona_bot/core/dotenv"
	"github.com/AnnonaOrg/annona_bot/core/features"
	"github.com/AnnonaOrg/annona_bot/core/log"
	"github.com/AnnonaOrg/osenv"
	tele "gopkg.in/telebot.v3"
)

func mainBot() {
	botToken := osenv.GetBotTelegramToken()
	botAPIProxyURL := osenv.GetBotTelegramAPIProxyURL()
	log.Debugf("GetBotTelegramAPIProxyURL(): %s", botAPIProxyURL)
	bot, err := tele.NewBot(tele.Settings{
		URL:    botAPIProxyURL,
		Token:  botToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	common.Must(err)

	features.Handle(bot)

	commands := []tele.Command{
		{
			Text:        "/start",
			Description: "开始",
		},
		{
			Text:        "/id",
			Description: "获取ID",
		},
		{
			Text:        "/ping",
			Description: "Ping",
		},
		// {
		// 	Text:        "/about",
		// 	Description: "About",
		// },
		{
			Text:        "/version",
			Description: "查看版本",
		},

		{
			Text:        "/reguser",
			Description: "注册登记",
		},
		{
			Text:        "/getuser",
			Description: "用户信息",
		},
		{
			Text:        "/addkeyword",
			Description: "添加关键词",
		},
		{
			Text:        "/delkeyworld",
			Description: "删除关键词",
		},
		{
			Text:        "/delallkeyworld",
			Description: "删除全部关键词",
		},
		{
			Text:        "/listkeyworld",
			Description: "列出关键词",
		},
		{
			Text:        "/renewuser",
			Description: "充值卡兑换",
		},
		{
			Text:        "/getcard",
			Description: "查询卡信息",
		},
	}
	bot.SetCommands(commands)

	bot.Start()
}
