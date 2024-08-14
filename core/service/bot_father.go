package service

import (
	"strings"

	"github.com/AnnonaOrg/annona_bot/core/constvar"

	"github.com/AnnonaOrg/annona_bot/core/log"
	"github.com/AnnonaOrg/annona_bot/core/utils"
	"github.com/AnnonaOrg/osenv"

	tele "gopkg.in/telebot.v3"
)

func SetBotFatherWebhook() {
	botToken := osenv.GetBotTelegramToken()
	// webhookURL := osenv.GetBotTelegramWebhookURL() //os.Getenv("BOT_TELEGRAM_WEBHOOK_URL")
	// if len(webhookURL) > 0 && strings.HasPrefix(webhookURL, "https") {
	// 	if tmpText, err := utils.SetTelegramWebhook(botToken, webhookURL+"/"+botToken); err != nil {
	// 		log.Errorf("SetTelegramWebhook(%s): %v", webhookURL, err)
	// 	} else {
	// 		log.Debugf("SetTelegramWebhook(%s): %s", webhookURL, tmpText)
	// 	}
	// } else {
	// 	log.Debugf("SetBotFatherWebhook(%s,%s)", botToken, webhookURL)
	// }
	// SetTelegramWebhook(botToken)
	SetBotWebhook(botToken)
}

// 初始化bot Commands 并设置webhook
func SetBotWebhook(botToken string) {
	bot, err := tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: true,
	})
	if err != nil {
		log.Errorf("SetBotWebhook(%s): %v", botToken, err)
		return
	}

	commands := constvar.Commands
	bot.SetCommands(commands)

	SetTelegramWebhook(botToken)
}

func SetTelegramWebhook(botToken string) {

	webhookURL := osenv.GetBotTelegramWebhookURL() //os.Getenv("BOT_TELEGRAM_WEBHOOK_URL")
	if len(webhookURL) > 0 && strings.HasPrefix(webhookURL, "https") {
		if tmpText, err := utils.SetTelegramWebhook(botToken, webhookURL+"/"+botToken); err != nil {
			log.Errorf("SetTelegramWebhook(%s): %v", webhookURL, err)
		} else {
			log.Debugf("SetTelegramWebhook(%s): %s", webhookURL, tmpText)
		}
	} else {
		log.Debugf("SetBotFatherWebhook(%s,%s)", botToken, webhookURL)
	}
}
