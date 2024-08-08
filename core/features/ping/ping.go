package ping

import (
	"fmt"

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
	text := fmt.Sprintf("Pong! %s%s \n@%s(%d)",
		c.Message().Sender.FirstName, c.Message().Sender.LastName,
		c.Message().Sender.Username, c.Message().Sender.ID,
	)

	// c.Delete()
	return c.Reply(text)
	// return c.Send(text)
}
