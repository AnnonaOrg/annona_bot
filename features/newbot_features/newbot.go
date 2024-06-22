package newbot_features

import (
	"fmt"
	"strings"
	"time"

	"github.com/AnnonaOrg/annona_bot/features"
	"github.com/AnnonaOrg/annona_bot/internal/constvar"
	"github.com/AnnonaOrg/annona_bot/internal/newbot_func"
	model_func "github.com/AnnonaOrg/annona_bot/internal/telebot_func"
	model "github.com/AnnonaOrg/annona_bot/model/telebot_info"
	"github.com/AnnonaOrg/osenv"

	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/newbot", OnNewbot)
}

// Command: /start <PAYLOAD>
func OnNewbot(c tele.Context) error {
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

	payload := c.Message().Payload
	payload = strings.TrimSpace(payload)
	if payload == "" || len(payload) < len("666666:aaaaaaaa_aaaaaaaaaaaaaaaaaaaaaaaaaaa") {
		return c.Reply("命令格式错误: 参考\n" + "```\n" + "/newbot botToken" + "\n```")
	}

	bot, err := newbot_func.NewBot(payload)
	if err != nil {
		return c.Reply("出了点问题: %v", err)
	}

	var item model.TeleBotInfo
	item.TelegramBotToken = bot.Token
	item.TelegramId = bot.Me.ID // bot.Me.ID
	item.TelegramUsername = bot.Me.Username
	item.TelegramUsernames = strings.Join(bot.Me.Usernames, ",")
	item.IsBot = bot.Me.IsBot
	item.IsForum = bot.Me.IsForum
	item.IsPremium = bot.Me.IsPremium
	item.FirstName = bot.Me.FirstName
	item.LastName = bot.Me.LastName
	item.ById = fmt.Sprintf("%d", c.Message().Sender.ID)

	if retText, err := model_func.DoAdd(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
		)
	} else {
		return c.Reply("ok " + retText)
	}
}
