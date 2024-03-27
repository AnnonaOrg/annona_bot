package user_info

import (
	"github.com/AnnonaOrg/annona_bot/model"
)

type UserInfo struct {
	model.Model

	InfoHash          string `json:"info_hash" form:"info_hash" gorm:"column:info_hash;not null;unique;"`
	AccoundPlatform   string `json:"accound_platform" form:"accound_platform" gorm:"column:accound_platform;not null;"`
	AccoundPlatformId string `json:"accound_platform_id" form:"accound_platform_id" gorm:"column:accound_platform_id;not null;"`

	TelegramChatId       int64  `json:"telegram_chat_id" form:"telegram_chat_id" gorm:"column:telegram_chat_id;"`
	TelegramUsername     string `json:"telegram_username" form:"telegram_username" gorm:"column:telegram_username;"` //@username
	TelegramChatUrl      string `json:"telegram_chat_url" form:"telegram_chat_url" gorm:"column:telegram_chat_url;"`
	TelegramStartBotId   int64  `json:"telegram_start_bot_id" form:"telegram_start_bot_id" gorm:"column:telegram_start_bot_id;"`       //激活机器人id
	TelegramNoticeChatId int64  `json:"telegram_notice_chat_id" form:"telegram_notice_chat_id" gorm:"column:telegram_notice_chat_id;"` //通知会话id

	TelegramFirstname string `json:"telegram_firstname" form:"telegram_firstname" gorm:"column:telegram_firstname;"`
	TelegramLasttname string `json:"telegram_lastname" form:"telegram_lastname" gorm:"column:telegram_lastname;"`
	// "exp" (expiration time)
	Exp int64 `json:"exp" form:"exp" gorm:"column:exp;not null;"`
	// 最后兑换卡
	LastCardUUID string `json:"last_card_uuid" form:"last_card_uuid" gorm:"column:last_card_uuid;"`
	// 最后签到时间
	LastSignDate string `json:"last_sign_date" form:"last_sign_date" gorm:"column:last_sign_date;"`
	// 邀请者
	Inviter string `json:"inviter" form:"inviter" gorm:"column:inviter;"`
	// 邀请码
	InviterCode string `json:"inviter_code" form:"inviter_code" gorm:"column:inviter_code;"`

	// 核验请求id
	ById string `json:"by_id" form:"by_id" gorm:"-"`

	TelegramStartBotUsername string `json:"-" form:"-" gorm:"-"`

	Page   int    `json:"-" form:"page" gorm:"-"`
	Size   int    `json:"-" form:"size" gorm:"-"`
	Filter string `json:"-" form:"filter" gorm:"-"`
}
