package setwebhook

import (
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/AnnonaOrg/annona_bot/common"
	_ "github.com/AnnonaOrg/annona_bot/main/distro/all"
	"github.com/AnnonaOrg/annona_bot/utils"
	tele "gopkg.in/telebot.v3"
)

// webhook=webhook+"/webhook/"+botToken
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}
	_, botToken, ok := strings.Cut(r.URL.Path, "setWebhook/")
	if len(botToken) <= 0 || !ok {
		log.Println("收到非法推送:(r.URL.Path)", r.URL.Path)
		return
	}
	if strings.Contains(botToken, "&") {
		if botTokenTmp, _, ok := strings.Cut(r.URL.Path, "&"); len(botTokenTmp) > 0 && ok {
			botToken = botTokenTmp
		}
	}

	bot, err := tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: true,
	})
	common.Must(err)

	commands := []tele.Command{
		{
			Text:        "/start",
			Description: "开始",
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
			Text:        "/reguser",
			Description: "注册登记",
		},
		{
			Text:        "/ping",
			Description: "Ping",
		},
		// {
		// 	Text:        "/about",
		// 	Description: "About",
		// },
	}
	bot.SetCommands(commands)

	webhookURL := os.Getenv("BOT_TELEGRAM_WEBHOOK_URL")
	if len(webhookURL) > 0 && strings.HasPrefix(webhookURL, "https") {
		utils.SetTelegramWebhook(botToken, webhookURL+"/"+botToken)
	}

}
