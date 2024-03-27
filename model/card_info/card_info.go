package card_info

import (
	"github.com/AnnonaOrg/annona_bot/model"
)

type CardInfo struct {
	model.Model

	CardHash string `json:"info_hash" form:"info_hash" gorm:"column:info_hash;not null;unique;"`
	CardUUID string `json:"card_uuid" form:"card_uuid" gorm:"column:card_uuid;not null;unique;"`
	// "exp" (expiration time)
	Exp int64 `json:"exp" form:"exp" gorm:"column:exp;not null;"`

	// 状态： 1 登记；2 待核销； 3 已核销卡;
	Stat int `json:"stat" form:"stat"  gorm:"column:stat;not null; "`
	//备注
	Note string `json:"note" form:"note" gorm:"column:note;"`

	// n天 30d 60d
	NDay string `json:"n_day" form:"n_day" gorm:"-"`
	// n张卡 30d 60d
	NNum int64 `json:"n_num" form:"n_num" gorm:"-"`

	// 核验请求id
	ById string `json:"by_id" form:"by_id" gorm:"-"`
	// 平台名称
	AccoundPlatform string `json:"accound_platform" form:"accound_platform" gorm:"-"`
	// 平台id
	AccoundPlatformId string `json:"accound_platform_id" form:"accound_platform_id" gorm:"-"`

	Page   int    `json:"-" form:"page" gorm:"-"`
	Size   int    `json:"-" form:"size" gorm:"-"`
	Filter string `json:"-" form:"filter" gorm:"-"`
}
