package tele_service

import (
	"fmt"
	"strconv"
	"strings"

	model_func "github.com/AnnonaOrg/annona_bot/core/blockformchatid_func"
	model_func2 "github.com/AnnonaOrg/annona_bot/core/blockformsenderid_func"
	"github.com/AnnonaOrg/annona_bot/core/blockword_func"
	"github.com/AnnonaOrg/annona_bot/core/constvar"

	"github.com/AnnonaOrg/annona_bot/core/keyword_func"
	"github.com/AnnonaOrg/annona_bot/core/log"
	"github.com/AnnonaOrg/annona_bot/core/service"
	model "github.com/AnnonaOrg/annona_bot/model/blockformchatid_info"
	model2 "github.com/AnnonaOrg/annona_bot/model/blockformsenderid_info"
	"github.com/AnnonaOrg/annona_bot/model/blockword_info"
	"github.com/AnnonaOrg/annona_bot/model/keyword_info"
	"github.com/AnnonaOrg/osenv"
	tele "gopkg.in/telebot.v3"
)

// 点击按钮回掉
func Callback(c tele.Context) {
	var payload string
	if callback := c.Callback(); callback != nil {
		payload = strings.TrimSpace(callback.Data)
	} else {
		return
	}

	log.Debugf("Callback().Data: %s", payload)
	payloadBody := ""
	if _, a, f := strings.Cut(payload, "|"); f && len(a) > 0 {
		payloadBody = a
	} else {
		c.Reply("请求异常: 未找到操作对象")
		return
	}
	senderID := c.Callback().Sender.ID
	ownerPlatform := osenv.GetPlatformType()

	switch {
	case strings.HasPrefix(payload, "/by_formsenderid"):
		{
			bySenderID, err := strconv.ParseInt(payloadBody, 10, 64)
			if err != nil {
				c.Reply(
					fmt.Sprintf("出了点问题: %v", payloadBody),
				)
				return
			}

			if retText, retTextAll, err := service.GetListKeyworldHistoryWithSenderID(bySenderID, 1); err != nil {
				c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				// return c.Reply(retText)
				if len(retTextAll) > 4096 {
					SendFileWithString(c, retTextAll, retText)
				} else {
					c.Reply(retText, tele.ModeHTML, tele.NoPreview)
				}
			}
			return
		}
	case strings.HasPrefix(payload, "/by_formkeyworld"):
		{
			if retText, retTextAll, err := service.GetListKeyworldHistoryWithKeyworld(payloadBody, 1); err != nil {
				c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				if len(retTextAll) > 4096 {
					SendFileWithString(c, retTextAll, retText)
				} else {
					c.Reply(retText, tele.ModeHTML, tele.NoPreview)
				}
			}
			return
		}
	case strings.HasPrefix(payload, "/block_formchatid"):
		{
			var item model.BlockformchatidInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := model_func.DoAdd(&item); err != nil {
				c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				c.Reply(retText)
			}
			return
		}

	case strings.HasPrefix(payload, "/block_formsenderid"):
		{
			var item model2.BlockformsenderidInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := model_func2.DoAdd(&item); err != nil {
				c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				c.Reply(retText)
			}
			return
		}

	case strings.HasPrefix(payload, "/add_keyword"):
		{
			var item keyword_info.KeyworldInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := keyword_func.DoAdd(&item); err != nil {
				c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				c.Reply(retText)
			}
			return
		}

	case strings.HasPrefix(payload, "/del_keyword"):
		{
			var item keyword_info.KeyworldInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := keyword_func.DoDel(&item); err != nil {
				c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				c.Reply(retText)
			}
			return
		}

	case strings.HasPrefix(payload, "/add_blockword"):
		{
			var item blockword_info.BlockworldInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := blockword_func.DoAdd(&item); err != nil {
				c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				c.Reply(retText)
			}
			return
		}

	case strings.HasPrefix(payload, "/del_blockword"):
		{
			var item blockword_info.BlockworldInfo
			item.OwnerPlatform = ownerPlatform
			item.OwnerChatId = senderID
			item.KeyWorld = payloadBody

			if retText, err := blockword_func.DoDel(&item); err != nil {
				c.Reply(
					fmt.Sprintf("出了点问题: %v", constvar.ERR_MSG_Server),
				)
			} else {
				c.Reply(retText)
			}
			return
		}
	default:
		return
	}

}
