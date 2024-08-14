package keyword_history_features

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/AnnonaOrg/annona_bot/core/service/tele_service"

	"github.com/AnnonaOrg/annona_bot/core/constvar"
	"github.com/AnnonaOrg/annona_bot/core/features"
	"github.com/AnnonaOrg/annona_bot/core/service"
	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/byid", OnByID)
	features.RegisterFeature("/byworld", OnByWorld)

}

// Command: /start <PAYLOAD>
func OnByID(c tele.Context) error {
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
				fmt.Sprintf("```\n%s\n```", "/byid 用户ID"),
			tele.ModeMarkdownV2,
		)
	}

	payload := strings.TrimSpace(c.Message().Payload)
	userID, err := strconv.ParseInt(payload, 10, 64)
	if err != nil {
		return c.Reply(
			"参考指令格式: \n"+
				fmt.Sprintf("```\n%s\n```", "/byid 用户ID"),
			tele.ModeMarkdownV2,
		)
	}
	retText, err := service.GetListKeyworldHistoryWithSenderID(userID, 1)
	if err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
		)
	}
	return c.Reply(retText)

}

// Command: /start <PAYLOAD>
func OnByWorld(c tele.Context) error {
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
				fmt.Sprintf("```\n%s\n```", "/byworld 关键词"),
			tele.ModeMarkdownV2,
		)
	}
	payload := strings.TrimSpace(c.Message().Payload)
	retText, retTextAll, err := service.GetListKeyworldHistoryWithKeyworld(payload, 1)
	if err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
		)
	}
	if len(retTextAll) > 4096 {
		tele_service.SendFileWithString(c, retTextAll, retText)
	} else {
		c.Reply(retText, tele.ModeHTML, tele.NoPreview)
	}
	return nil
}
