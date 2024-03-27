package blockword_info

import (
	"github.com/AnnonaOrg/annona_bot/model"
)

type BlockworldInfo struct {
	model.Model

	InfoHash      string `json:"info_hash" form:"info_hash" gorm:"column:info_hash;not null;"`
	OwnerInfoHash string `json:"owner_info_hash" form:"owner_info_hash" gorm:"column:owner_info_hash;not null;"`
	OwnerPlatform string `json:"owner_platform" form:"owner_platform" gorm:"column:owner_platform;not null;"`
	OwnerChatId   int64  `json:"owner_chat_id" form:"owner_chat_id" gorm:"column:owner_chat_id;not null;"`
	SearchChatId  int64  `json:"search_chat_id" form:"search_chat_id" gorm:"column:search_chat_id;"`
	KeyWorld      string `json:"key_world" form:"key_world" gorm:"column:key_world;"`

	// 核验请求id
	ById string `json:"by_id" form:"by_id" gorm:"-"`

	Page   int    `json:"-" form:"page" gorm:"-"`
	Size   int    `json:"-" form:"size" gorm:"-"`
	Filter string `json:"-" form:"filter" gorm:"-"`
}
