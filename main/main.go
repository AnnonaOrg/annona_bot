package main

import (
	// _ "github.com/AnnonaOrg/annona_bot/api"
	// _ "github.com/AnnonaOrg/annona_bot/api/push"
	// _ "github.com/AnnonaOrg/annona_bot/api/setwebhook"

	"time"

	"github.com/AnnonaOrg/annona_bot/common"
	"github.com/AnnonaOrg/annona_bot/features"
	_ "github.com/AnnonaOrg/annona_bot/internal/dotenv"
	_ "github.com/AnnonaOrg/annona_bot/main/distro/all"
	"github.com/AnnonaOrg/osenv"

	tele "gopkg.in/telebot.v3"
)

func main() {
	botToken := osenv.GetBotTelegramToken()
	bot, err := tele.NewBot(tele.Settings{
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
		{
			Text:        "/ping",
			Description: "Ping",
		},
		{
			Text:        "/about",
			Description: "About",
		},
	}
	bot.SetCommands(commands)

	bot.Start()
}
