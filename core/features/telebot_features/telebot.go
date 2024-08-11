package telebot_features

import (
	"fmt"
	"strings"
	"time"

	"github.com/AnnonaOrg/annona_bot/core/constvar"
	"github.com/AnnonaOrg/annona_bot/core/features"
	model_func "github.com/AnnonaOrg/annona_bot/core/telebot_func"
	model "github.com/AnnonaOrg/annona_bot/model/telebot_info"
	"github.com/AnnonaOrg/osenv"
	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/botenable", OnEnableTelebotInfo)
}

// 启用机器人
// Command: /start <PAYLOAD>
func OnEnableTelebotInfo(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}
	if !osenv.IsBotManagerID(c.Message().Sender.ID) {
		return c.Reply("🈲️非法指令")
	}

	var item model.TeleBotInfo
	item.TelegramBotToken = c.Bot().Token
	item.TelegramId = c.Bot().Me.ID // bot.Me.ID
	item.TelegramUsername = c.Bot().Me.Username
	item.TelegramUsernames = strings.Join(c.Bot().Me.Usernames, ",")
	item.IsBot = c.Bot().Me.IsBot
	item.IsForum = c.Bot().Me.IsForum
	item.IsPremium = c.Bot().Me.IsPremium
	item.FirstName = c.Bot().Me.FirstName
	item.LastName = c.Bot().Me.LastName
	item.ById = fmt.Sprintf("%d", c.Message().Sender.ID)

	if retText, err := model_func.DoAdd(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
		)
	} else {
		if len(retText) == 0 {
			retText = "👌"
		}
		return c.Reply(retText)
	}
}
