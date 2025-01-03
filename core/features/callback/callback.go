package callback

import (
	"github.com/AnnonaOrg/annona_bot/core/features"
	"github.com/AnnonaOrg/annona_bot/core/service/tele_service"

	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature(tele.OnCallback, OnCallback)
}

func OnCallback(c tele.Context) error {
	tele_service.Callback(c)
	return nil
}
