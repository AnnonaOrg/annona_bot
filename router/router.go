package router

import (
	"github.com/AnnonaOrg/annona_bot/handler/api_handler"
	"github.com/AnnonaOrg/annona_bot/handler/bot_handler"
	"github.com/AnnonaOrg/annona_bot/handler/webhook_handler/tele_handler"
	"github.com/AnnonaOrg/annona_bot/handler/wspush"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())

	g.Use(mw...)
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	g.MaxMultipartMemory = 8 << 20 // 8 MiB

	g.NoRoute(api_handler.ApiNotFound)
	g.GET("/", api_handler.ApiHello)
	g.GET("/ping", api_handler.ApiPing)

	webhookR := g.Group("/webhook")
	{
		webhookR.POST("/tele/:botToken", tele_handler.Update)

		webhookR.POST("/set/:botToken", bot_handler.SetWebhook)
	}

	wspushR := g.Group("/ws")
	{
		wspushR.POST("/push_v2/:channel", wspush.WSPush)
	}

	return g
}
