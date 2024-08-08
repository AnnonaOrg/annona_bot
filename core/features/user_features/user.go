package user_features

import (
	"fmt"
	"strconv"
	"time"

	"github.com/AnnonaOrg/annona_bot/core/constvar"
	"github.com/AnnonaOrg/annona_bot/core/features"
	model_func "github.com/AnnonaOrg/annona_bot/core/user_func"
	model "github.com/AnnonaOrg/annona_bot/model/user_info"
	"github.com/AnnonaOrg/osenv"
	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/reguser", OnRegUserInfo)
	features.RegisterFeature("/getuser", OnGetUserInfo)
	features.RegisterFeature("/renewuser", OnRenewUserInfo)
	features.RegisterFeature("/updatenoticechatid", OnUpdateNoticeChatId)
	// 签到
	features.RegisterFeature("/sign", OnQianDao)
}

// 用户注册
// Command: /start <PAYLOAD>
func OnRegUserInfo(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}

	helloStr := fmt.Sprintf("Hello! %s(%d)", c.Message().Sender.Username, c.Message().Sender.ID)

	var item model.UserInfo
	item.AccoundPlatform = osenv.GetPlatformType() // constvar.PLATFORM_TYPE_TELE
	item.AccoundPlatformId = fmt.Sprintf("%d", c.Message().Sender.ID)
	item.TelegramChatId = c.Message().Sender.ID
	item.TelegramUsername = c.Message().Sender.Username
	item.TelegramStartBotId = c.Bot().Me.ID
	item.TelegramFirstname = c.Message().Sender.FirstName
	item.TelegramLasttname = c.Message().Sender.LastName

	if _, err := model_func.DoAdd(&item); err != nil {
		helloStr = fmt.Sprintf("%s \n出了点问题: %v", helloStr, constvar.ERR_MSG_Server)
	}
	return c.Reply(helloStr)
}

// 用户信息
// Command: /start <PAYLOAD>
func OnGetUserInfo(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Printf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}

	helloStr := fmt.Sprintf(
		"Hello! %s %s \n%s %d \n",
		c.Message().Sender.FirstName, c.Message().Sender.LastName,
		c.Message().Sender.Username, c.Message().Sender.ID,
	)

	var item model.UserInfo
	item.AccoundPlatform = osenv.GetPlatformType() //  constvar.PLATFORM_TYPE_TELE
	item.AccoundPlatformId = fmt.Sprintf("%d", c.Message().Sender.ID)
	item.TelegramChatId = c.Message().Sender.ID
	item.TelegramUsername = c.Message().Sender.Username
	item.TelegramStartBotId = c.Bot().Me.ID
	item.TelegramStartBotUsername = c.Bot().Me.Username
	item.TelegramFirstname = c.Message().Sender.FirstName
	item.TelegramLasttname = c.Message().Sender.LastName

	if retText, err := model_func.DoGet(&item); err != nil {
		return c.Reply(
			helloStr + fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
		)
	} else {
		return c.Reply(
			helloStr+retText,
			tele.NoPreview,
		)
	}

}

// 兑换卡密
// Command: /start <PAYLOAD>
func OnRenewUserInfo(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}
	if len(c.Message().Payload) == 0 {
		return c.Reply(
			"参考指令格式: \n"+
				fmt.Sprintf("```\n%s\n```", "/renewuser 卡ID"),
			tele.ModeMarkdownV2,
		)
	}

	var item model.UserInfo
	item.AccoundPlatform = osenv.GetPlatformType() //  constvar.PLATFORM_TYPE_TELE
	item.AccoundPlatformId = fmt.Sprintf("%d", c.Message().Sender.ID)
	item.TelegramChatId = c.Message().Sender.ID
	item.TelegramUsername = c.Message().Sender.Username
	item.TelegramStartBotId = c.Bot().Me.ID
	item.TelegramFirstname = c.Message().Sender.FirstName
	item.TelegramLasttname = c.Message().Sender.LastName

	item.LastCardUUID = c.Message().Payload

	if retText, err := model_func.DoRenew(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
		)
	} else {
		return c.Reply(retText)
	}
}

// 签到
// Command: /start <PAYLOAD>
func OnQianDao(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}
	if len(c.Message().Payload) == 0 {
		return c.Reply(
			"参考指令格式: \n"+
				fmt.Sprintf("```\n%s\n```", "/sign 祝福语"),
			tele.ModeMarkdownV2,
		)
	}

	var item model.UserInfo
	item.AccoundPlatform = osenv.GetPlatformType() //  constvar.PLATFORM_TYPE_TELE
	item.AccoundPlatformId = fmt.Sprintf("%d", c.Message().Sender.ID)
	item.TelegramChatId = c.Message().Sender.ID
	item.TelegramUsername = c.Message().Sender.Username
	item.TelegramStartBotId = c.Bot().Me.ID
	item.TelegramFirstname = c.Message().Sender.FirstName
	item.TelegramLasttname = c.Message().Sender.LastName

	item.LastCardUUID = c.Message().Payload

	if retText, err := model_func.DoSign(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
		)
	} else {
		return c.Reply(retText)
	}
}

// 更新通知会话id
// Command: /start <PAYLOAD>
func OnUpdateNoticeChatId(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}
	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}

	if len(c.Message().Payload) == 0 {
		return c.Reply(
			"参考指令格式: \n"+
				fmt.Sprintf("```\n%s\n```", "/updatenoticechatid ChatID"),
			tele.ModeMarkdownV2,
		)
	}

	var item model.UserInfo
	item.AccoundPlatform = osenv.GetPlatformType() //  constvar.PLATFORM_TYPE_TELE
	item.AccoundPlatformId = fmt.Sprintf("%d", c.Message().Sender.ID)
	item.TelegramChatId = c.Message().Sender.ID
	item.TelegramUsername = c.Message().Sender.Username
	item.TelegramStartBotId = c.Bot().Me.ID
	item.TelegramFirstname = c.Message().Sender.FirstName
	item.TelegramLasttname = c.Message().Sender.LastName

	noticeChatId, err := strconv.ParseInt(c.Message().Payload, 10, 64)
	if err != nil || noticeChatId == 0 {
		return c.Reply("ChatID错误")
	}
	item.TelegramNoticeChatId = noticeChatId

	if retText, err := model_func.DoUpdateNoticeChatId(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
		)
	} else {
		return c.Reply(retText)
	}
}
