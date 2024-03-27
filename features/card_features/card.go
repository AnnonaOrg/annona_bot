package card_features

import (
	"fmt"
	"strconv"

	"github.com/AnnonaOrg/annona_bot/features"
	model_func "github.com/AnnonaOrg/annona_bot/internal/card_func"
	model "github.com/AnnonaOrg/annona_bot/model/card_info"
	"github.com/AnnonaOrg/osenv"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/convert", OnConvertCardInfo)
	features.RegisterFeature("/addcard", OnAddCardInfo)
	features.RegisterFeature("/getcard", OnGetCardInfo)
}

// Command: /start <PAYLOAD>
func OnConvertCardInfo(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	// if !osenv.IsBotManagerID(c.Message().Sender.ID) {
	// 	return c.Reply("🈲️非法指令")
	// }

	cmdFmtErr := func() error {
		return c.Reply(
			"指令格式错误，参考: \n"+
				fmt.Sprintf("```\n%s\n```", "/convert nNum nDay"),
			tele.ModeMarkdownV2,
		)
	}

	var item model.CardInfo
	item.ById = fmt.Sprintf("%d", c.Message().Sender.ID)
	item.AccoundPlatform = osenv.GetPlatformType()
	item.AccoundPlatformId = fmt.Sprintf("%d", c.Message().Sender.ID)

	args := c.Args()
	switch len(args) {
	case 1:
		if n, err := strconv.ParseInt(c.Message().Payload, 10, 64); err != nil {
			return cmdFmtErr()
		} else {
			if n > 0 {
				item.NNum = n
			} else {
				item.NNum = 1
			}
		}

	case 2:
		nNum := args[0]
		nDay := args[1]
		if n, err := strconv.ParseInt(nNum, 10, 64); err != nil {
			return cmdFmtErr()
		} else {
			if n > 0 {
				item.NNum = n
			} else {
				item.NNum = 1
			}
		}
		item.NDay = nDay

	default:
		return cmdFmtErr()
	}

	if retText, err := model_func.DoConvert(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", err),
		)
	} else {
		return c.Reply(retText)
	}
}

// Command: /start <PAYLOAD>
func OnAddCardInfo(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	if !osenv.IsBotManagerID(c.Message().Sender.ID) {
		return c.Reply("🈲️非法指令")
	}

	cmdFmtErr := func() error {
		return c.Reply(
			"指令格式错误，参考: \n"+
				fmt.Sprintf("```\n%s\n```", "/addcard nNum nDay"),
			tele.ModeMarkdownV2,
		)
	}

	var item model.CardInfo
	item.ById = fmt.Sprintf("%d", c.Message().Sender.ID)
	item.AccoundPlatform = osenv.GetPlatformType()
	item.AccoundPlatformId = fmt.Sprintf("%d", c.Message().Sender.ID)

	args := c.Args()
	switch len(args) {
	case 1:
		if n, err := strconv.ParseInt(c.Message().Payload, 10, 64); err != nil {
			return cmdFmtErr()
		} else {
			if n > 0 {
				item.NNum = n
			} else {
				item.NNum = 1
			}
		}

	case 2:
		nNum := args[0]
		nDay := args[1]
		if n, err := strconv.ParseInt(nNum, 10, 64); err != nil {
			return cmdFmtErr()
		} else {
			if n > 0 {
				item.NNum = n
			} else {
				item.NNum = 1
			}
		}
		item.NDay = nDay

	default:
		return cmdFmtErr()
	}

	if retText, err := model_func.DoAdd(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", err),
		)
	} else {
		return c.Reply(retText)
	}
}

// Command: /start <PAYLOAD>
func OnGetCardInfo(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}
	if len(c.Message().Payload) == 0 {
		return c.Reply(
			"参考指令格式: \n"+
				fmt.Sprintf("```\n%s\n```", "/getcard 卡ID"),
			tele.ModeMarkdownV2,
		)
	}

	var item model.CardInfo
	item.CardUUID = c.Message().Payload

	if retText, err := model_func.DoGet(&item); err != nil {
		return c.Reply(
			fmt.Sprintf("出了点问题: %v", err),
		)
	} else {
		return c.Reply(retText)
	}
}
