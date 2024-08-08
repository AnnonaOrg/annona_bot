package tele_handler

import (
	"encoding/json"

	_ "github.com/AnnonaOrg/annona_bot/cmd/annona_bot/distro/all"
	"github.com/AnnonaOrg/annona_bot/core/features"
	"github.com/AnnonaOrg/annona_bot/core/log"
	"github.com/gin-gonic/gin"
	tele "gopkg.in/telebot.v3"
)

// webhook=webhook+"/webhook/"+botToken
func Update(c *gin.Context) {
	botToken := c.Param("botToken")
	if len(botToken) == 0 {
		log.Debugf("收到非法推送: %s", c.Request.URL)
		return
	}
	requestBody, err := c.GetRawData()
	if err != nil {
		log.Errorf("GetRawData(): %v", err)
		return
	}
	log.Debugf("requestBody: %s", string(requestBody))

	bot, err := tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: true,
	})
	if err != nil {
		log.Errorf("NewBot出错: %v", err)
		return
	}

	features.Handle(bot)

	var u tele.Update
	err = json.Unmarshal(requestBody, &u)
	if err != nil {
		log.Errorf("json.Unmarshal(%s, &tele.Update): %v", string(requestBody), err)
		return
	}
	bot.ProcessUpdate(u)
}
