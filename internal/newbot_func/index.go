package newbot_func

import (
	"fmt"

	"strings"

	"github.com/AnnonaOrg/osenv"

	"github.com/AnnonaOrg/annona_bot/utils"
	tele "gopkg.in/telebot.v3"
)

// webhook=webhook+"/webhook/"+botToken
func NewBot(botToken string) (*tele.Bot, error) {
	bot, err := tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: true,
	})
	if err != nil {
		return nil, err
	}
	// common.Must(err)

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
			Text:        "/reguser",
			Description: "注册登记",
		},
		{
			Text:        "/updatenoticechatid",
			Description: "更新通知会话ID",
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
			Text:        "/ping",
			Description: "Ping",
		},
		{
			Text:        "/version",
			Description: "查看版本号",
		},
	}
	bot.SetCommands(commands)

	webhookURL := osenv.GetBotTelegramWebhookURL() //os.Getenv("BOT_TELEGRAM_WEBHOOK_URL")
	if len(webhookURL) > 0 && strings.HasPrefix(webhookURL, "https") {
		if _, err := utils.SetTelegramWebhook(botToken, webhookURL+"/"+botToken); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("BOT_TELEGRAM_WEBHOOK_URL is NULL")
	}
	return bot, nil
}
