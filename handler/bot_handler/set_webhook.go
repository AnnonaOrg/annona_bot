package bot_handler

import (
	"github.com/AnnonaOrg/annona_bot/core/service"
	"github.com/AnnonaOrg/annona_bot/handler"
	"github.com/AnnonaOrg/pkg/errno"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// /ws/push_v2/channelStr
func SetWebhook(c *gin.Context) {
	botToken := c.Param("botToken")
	if len(botToken) == 0 {
		log.Debugf("收到非法推送: %s", c.Request.URL.Path)
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	// log.Debugf("Request.URL: %s", c.Request.URL.String())
	handler.SendResponse(c, nil, "ok")
	go service.SetBotWebhook(botToken)
}
