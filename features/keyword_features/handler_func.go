package keyword_features

import (
	"fmt"
	"time"

	"github.com/AnnonaOrg/osenv"

	model_func "github.com/AnnonaOrg/annona_bot/internal/keyword_func"
	model "github.com/AnnonaOrg/annona_bot/model/keyword_info"

	"github.com/AnnonaOrg/annona_bot/features"
	"github.com/AnnonaOrg/annona_bot/internal/constvar"

	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/addkeyword", OnAdd)
	features.RegisterFeature("/delkeyworld", OnDel)
	features.RegisterFeature("/delallkeyworld", OnDelall)
	features.RegisterFeature("/listkeyworld", OnList)
}

// Command: /start <PAYLOAD>
func OnAdd(c tele.Context) error {
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
				fmt.Sprintf("```\n%s\n```", "/addkeyword 关键词"),
			tele.ModeMarkdownV2,
		)
	}

	var item model.KeyworldInfo
	item.OwnerPlatform = osenv.GetPlatformType() //  constvar.PLATFORM_TYPE_TELE
	item.OwnerChatId = c.Message().Sender.ID
	item.KeyWorld = c.Message().Payload

	if retText, err := model_func.DoAdd(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", err),
		)
	} else {
		return c.Reply(retText)
	}
}

// Command: /start <PAYLOAD>
func OnDel(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return c.Reply("请重试!")
	}
	if len(c.Message().Payload) == 0 {
		return c.Reply(
			"参考指令格式: \n"+
				fmt.Sprintf("```\n%s\n```", "/delkeyworld 关键词"),
			tele.ModeMarkdownV2,
		)
	}

	var item model.KeyworldInfo
	item.OwnerPlatform = osenv.GetPlatformType() //  constvar.PLATFORM_TYPE_TELE
	item.OwnerChatId = c.Message().Sender.ID
	item.KeyWorld = c.Message().Payload

	if retText, err := model_func.DoDel(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", err),
		)
	} else {
		return c.Reply(retText)
	}
}

// Command: /start <PAYLOAD>
func OnDelall(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}

	var item model.KeyworldInfo
	item.OwnerPlatform = osenv.GetPlatformType() //  constvar.PLATFORM_TYPE_TELE
	item.OwnerChatId = c.Message().Sender.ID
	// item.KeyWorld = c.Message().Payload

	if retText, err := model_func.DoDelall(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", err),
		)
	} else {
		return c.Reply(retText)
	}
}

// Command: /start <PAYLOAD>
func OnList(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}

	var item model.KeyworldInfo
	item.OwnerPlatform = osenv.GetPlatformType() //  constvar.PLATFORM_TYPE_TELE
	item.OwnerChatId = c.Message().Sender.ID
	// item.KeyWorld = c.Message().Payload

	if retText, err := model_func.DoList(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", err),
		)
	} else {
		return c.Reply(retText)
	}
}
