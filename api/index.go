package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	// "github.com/AnnonaOrg/annona_bot/common"
	"github.com/AnnonaOrg/annona_bot/features"
	_ "github.com/AnnonaOrg/annona_bot/main/distro/all"

	tele "gopkg.in/telebot.v3"
)

// webhook=webhook+"/webhook/"+botToken
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// log.Debugf("收到非法请求: %s", r.URL.Path)
		return
	}
	_, botToken, ok := strings.Cut(r.URL.Path, "webhook/")
	if len(botToken) == 0 || !ok {
		log.Debugf("收到非法推送: %s", r.URL.Path)
		return
	}
	body, err := io.ReadAll(r.Body)
	// common.Must(err)
	if err != nil {
		log.Errorf("io.ReadAll(r.Body): %v", err)
		return
	}
	log.Debugf("body: %s", string(body))

	bot, err := tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: true,
	})
	if err != nil {
		log.Errorf("NewBot出错: %v", err)
		return
	}
	// common.Must(err)
	// commands := []tele.Command{
	// 	{
	// 		Text:        "/ping",
	// 		Description: "Ping",
	// 	},
	// 	{
	// 		Text:        "/start",
	// 		Description: "Start",
	// 	},
	// }
	// bot.SetCommands(commands)
	// webhookURL := os.Getenv("BOT_TELEGRAM_WEBHOOK_URL")
	// if len(webhookURL) > 0 && strings.HasPrefix(webhookURL, "https") {
	// 	utils.SetTelegramWebhook(botToken, webhookURL+"/"+botToken)
	// }
	features.Handle(bot)

	var u tele.Update
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Errorf("json.Unmarshal(%s, &tele.Update): %v", string(body), err)
		return
	}
	// common.Must(err)
	bot.ProcessUpdate(u)
}
