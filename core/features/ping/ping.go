package ping

import (
	"github.com/AnnonaOrg/annona_bot/core/features"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/ping", OnPing)
}

func OnPing(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	return c.Reply("Pong!")
}
