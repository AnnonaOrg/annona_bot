package about

import (
	"github.com/AnnonaOrg/annona_bot/core/constvar"
	"github.com/AnnonaOrg/annona_bot/core/features"
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
	text := constvar.APPAbout()
	return c.Reply(text)
}
func OnVersion(c tele.Context) error {
	text := constvar.APPVersion()
	return c.Reply(text)
}
