package wspush

import (
	"io"

	"github.com/AnnonaOrg/annona_bot/handler"
	"github.com/AnnonaOrg/pkg/errno"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// /ws/push_v2/channelStr
func WSPush(c *gin.Context) {
	channelStr := c.Param("channel")
	if len(channelStr) == 0 {
		log.Debugf("收到非法推送: %s", c.Request.URL.Path)
		return
	}
	log.Debugf("Request.URL: %s", c.Request.URL.String())

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, "")
		log.Errorf("io.ReadAll(r.Body): %v", err)
		return
	}

	if err := PushMsgData(body); err != nil {
		handler.SendResponse(c, errno.InternalServerError, "")
		log.Errorf("PushMsgData(%s): %v", string(body), err)
		return
	}
	handler.SendResponse(c, nil, "ok")
}
