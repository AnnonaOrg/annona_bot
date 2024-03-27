package about

import (
	"github.com/AnnonaOrg/annona_bot/features"
	"github.com/AnnonaOrg/annona_bot/internal/constvar"
	"github.com/AnnonaOrg/osenv"

	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/about", OnProcess)
	features.RegisterFeature("/version", OnVersion)
}
func OnProcess(c tele.Context) error {
	if !osenv.IsBotManagerID(c.Message().Sender.ID) {
		return nil
	}
	text := constvar.About()
	return c.Reply(text)
}
func OnVersion(c tele.Context) error {
	text := constvar.Version()
	return c.Reply(text)
}
