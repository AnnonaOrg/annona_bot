package callback

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/AnnonaOrg/annona_bot/core/blockword_func"
	"github.com/AnnonaOrg/annona_bot/core/keyword_func"
	"github.com/AnnonaOrg/annona_bot/core/log"
	"github.com/AnnonaOrg/annona_bot/core/service"
	"github.com/AnnonaOrg/annona_bot/model/blockword_info"
	"github.com/AnnonaOrg/annona_bot/model/keyword_info"

	model_func "github.com/AnnonaOrg/annona_bot/core/blockformchatid_func"
	model "github.com/AnnonaOrg/annona_bot/model/blockformchatid_info"

	model_func2 "github.com/AnnonaOrg/annona_bot/core/blockformsenderid_func"
	model2 "github.com/AnnonaOrg/annona_bot/model/blockformsenderid_info"

	"github.com/AnnonaOrg/annona_bot/core/constvar"
	"github.com/AnnonaOrg/annona_bot/core/features"
	"github.com/AnnonaOrg/osenv"

	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature(tele.OnCallback, OnCallback)
}

// 点击按钮回掉
func OnCallback(c tele.Context) error {
	// if !c.Message().Private() {
	// 	return nil
	// }

	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}

	payload := c.Callback().Data
	payload = strings.TrimSpace(payload)
	log.Debugf("Callback().Data: %s", payload)
	payloadBody := ""
	if _, a, f := strings.Cut(payload, "|"); f && len(a) > 0 {
		payloadBody = a
	} else {
		return c.Reply("请求异常: 未找到操作对象")
	}
	senderID := c.Callback().Sender.ID
	ownerPlatform := osenv.GetPlatformType()

	// return c.Reply("payloadBody: " + payloadBody)
	// block_formsenderid
	switch {
	case strings.HasPrefix(payload, "/by_formsenderid"):
		{
			bySenderID, err := strconv.ParseInt(payloadBody, 10, 64)
			if err != nil {
				return c.Reply(
					fmt.Sprintf("出了点问题: %v", payloadBody),
				)
			}

			if retText, err := service.GetListKeyworldHistoryWithSenderID(bySenderID, 1); err != nil {
				return c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				return c.Reply(retText)
			}
		}
	case strings.HasPrefix(payload, "/by_formkeyworld"):
		{
			if retText, err := service.GetListKeyworldHistoryWithKeyworld(payloadBody, 1); err != nil {
				return c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				return c.Reply(retText, tele.ModeHTML, tele.NoPreview)
			}
		}
	case strings.HasPrefix(payload, "/block_formchatid"):
		{
			var item model.BlockformchatidInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := model_func.DoAdd(&item); err != nil {
				return c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				return c.Reply(retText)
			}
		}

	case strings.HasPrefix(payload, "/block_formsenderid"):
		{
			var item model2.BlockformsenderidInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := model_func2.DoAdd(&item); err != nil {
				return c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				return c.Reply(retText)
			}
		}

	case strings.HasPrefix(payload, "/add_keyword"):
		{
			var item keyword_info.KeyworldInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := keyword_func.DoAdd(&item); err != nil {
				return c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				return c.Reply(retText)
			}
		}

	case strings.HasPrefix(payload, "/del_keyword"):
		{
			var item keyword_info.KeyworldInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := keyword_func.DoDel(&item); err != nil {
				return c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				return c.Reply(retText)
			}
		}

	case strings.HasPrefix(payload, "/add_blockword"):
		{
			var item blockword_info.BlockworldInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := blockword_func.DoAdd(&item); err != nil {
				return c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				return c.Reply(retText)
			}
		}

	case strings.HasPrefix(payload, "/del_blockword"):
		{
			var item blockword_info.BlockworldInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := blockword_func.DoDel(&item); err != nil {
				return c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				return c.Reply(retText)
			}
		}
	default:
		return nil
	}

}
