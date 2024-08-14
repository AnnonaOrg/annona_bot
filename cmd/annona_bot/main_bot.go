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
			Text:        "/sign",
			Description: "签到",
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
			Text:        "/getuser",
			Description: "用户信息",
		},

		{
			Text:        "/updatenoticechatid",
			Description: "更换通知会话ID",
		},
		{
			Text:        "/convert",
			Description: "兑换克莱因瓶",
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
			Text:        "/addblockword",
			Description: "添加屏蔽关键词",
		},
		{
			Text:        "/delblockword",
			Description: "删除屏蔽关键词",
		},
		{
			Text:        "/delallblockword",
			Description: "删除全部屏蔽关键词",
		},
		{
			Text:        "/listblockword",
			Description: "列出屏蔽关键词",
		},

		{
			Text:        "/addblockformsenderid",
			Description: "添加屏蔽发送者ID",
		},
		{
			Text:        "/delblockformsenderid",
			Description: "删除屏蔽发送者ID",
		},
		{
			Text:        "/delallblockformsenderid",
			Description: "删除全部屏蔽发送者ID",
		},
		{
			Text:        "/listblockformsenderid",
			Description: "列出屏蔽发送者ID",
		},

		{
			Text:        "/addblockformchatid",
			Description: "添加屏蔽群组ID",
		},
		{
			Text:        "/delblockformchatid",
			Description: "删除屏蔽群组ID",
		},
		{
			Text:        "/delallblockformchatid",
			Description: "删除全部屏蔽群组ID",
		},
		{
			Text:        "/listblockformchatid",
			Description: "列出屏蔽群组ID",
		},
		{
			Text:        "/byid",
			Description: "通过用户ID查询",
		},
		{
			Text:        "/byworld",
			Description: "通过关键词查询",
		},
		{
			Text:        "/reguser",
			Description: "注册登记",
		},
		{
			Text:        "/ping",
			Description: "Ping",
		},
		{
			Text:        "/version",
			Description: "查看版本",
		},
	}
	bot.SetCommands(commands)

	bot.Start()
}
