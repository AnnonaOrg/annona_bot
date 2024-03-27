package text

import (
	"time"

	"github.com/AnnonaOrg/annona_bot/features"
	"github.com/AnnonaOrg/annona_bot/internal/constvar"
	"github.com/AnnonaOrg/annona_bot/utils"
	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature(tele.OnText, OnTextEx)
}

func OnTextEx(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}
	if (time.Now().Unix() - c.Message().Unixtime) > constvar.TIME_OUT_MAX_seconds {
		log.Debugf("忽略超过 %d s(秒) 的消息处理", constvar.TIME_OUT_MAX_seconds)
		return nil
	}

	text := utils.GetStringRuneN(c.Message().Text, 20)
	if text == "" {
		return nil
	}

	selector := &tele.ReplyMarkup{}
	btnAddKeyword := selector.Data("添加监控词", "/add_keyword", text)
	btnDelKeyword := selector.Data("删除监控词", "/del_keyword", text)
	btnAddBlockword := selector.Data("添加屏蔽词", "/add_blockword", text)
	btnDelBlockword := selector.Data("删除屏蔽词", "/del_blockword", text)
	btnLink := selector.URL("购买充值卡🛒", "https://t.me/annonaCardBot")
	btnLinkServiceSupport := selector.URL("支持频道✅", "https://t.me/annonaOrg")
	btnLinkSubmitNewGroup := selector.URL("提交群组📨", "https://t.me/annonaGroup")
	selector.Inline(
		selector.Row(btnAddKeyword, btnDelKeyword),
		selector.Row(btnAddBlockword, btnDelBlockword),
		selector.Row(btnLinkServiceSupport, btnLinkSubmitNewGroup),
		selector.Row(btnLink),
	)

	c.Reply(text, selector)

	return nil
}
