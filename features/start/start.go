package start

import (
	"fmt"
	"time"

	"github.com/AnnonaOrg/annona_bot/features"
	"github.com/AnnonaOrg/annona_bot/internal/constvar"
	model_func "github.com/AnnonaOrg/annona_bot/internal/user_func"
	model "github.com/AnnonaOrg/annona_bot/model/user_info"
	"github.com/AnnonaOrg/osenv"
	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/start", Onstart)
}

// Command: /start <PAYLOAD>
func Onstart(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}
	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}

	helloStr := fmt.Sprintf("Hello! %s %s\n@%s (%d)",
		c.Message().Sender.FirstName, c.Message().Sender.LastName,
		c.Message().Sender.Username, c.Message().Sender.ID,
	)

	var item model.UserInfo
	item.AccoundPlatform = osenv.GetPlatformType() // "tele" //constvar.PLATFORM_TYPE_TELE
	item.AccoundPlatformId = fmt.Sprintf("%d", c.Message().Sender.ID)
	item.TelegramChatId = c.Message().Sender.ID
	item.TelegramUsername = c.Message().Sender.Username
	item.TelegramStartBotId = c.Bot().Me.ID
	item.TelegramFirstname = c.Message().Sender.FirstName
	item.TelegramLasttname = c.Message().Sender.LastName

	payload := c.Message().Payload
	if len(payload) > 0 {
		item.Inviter = payload
		helloStr = helloStr + "\n" +
			fmt.Sprintf("邀请者ID: "+payload)
	}

	if _, err := model_func.DoAdd(&item); err != nil {
		helloStr = fmt.Sprintf("%s \n出了点小问题: %v", helloStr, constvar.ERR_MSG_Server)
		if osenv.IsBotManagerID(c.Message().Sender.ID) {
			helloStr = fmt.Sprintf("%s \n出了点小问题: %v", helloStr, err)
		}
	}

	if osenv.IsBotManagerID(c.Message().Sender.ID) {
		helloStr = helloStr + "\n" +
			"专属指令: " + "\n" +
			"/about" + "\n" +
			"/addcard nNum nDay" + "\n" +
			"/botenable" + "\n" +
			"/newbot botToken"
	}
	return c.Reply(helloStr)
}
