package main

import (
	"time"

	_ "github.com/AnnonaOrg/annona_bot/cmd/annona_bot/distro/all"
	"github.com/AnnonaOrg/annona_bot/common"
	"github.com/AnnonaOrg/annona_bot/core/constvar"
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

	commands := constvar.Commands
	bot.SetCommands(commands)

	bot.Start()
}
