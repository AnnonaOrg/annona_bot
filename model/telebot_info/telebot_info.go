package telebot_info

import (
	"github.com/AnnonaOrg/annona_bot/model"
)

type TeleBotInfo struct {
	model.Model

	TelegramId        int64  `json:"telegram_id" form:"telegram_id" gorm:"column:telegram_id;not null;"`
	TelegramUsername  string `json:"telegram_username" form:"telegram_username" gorm:"column:telegram_username;"`
	TelegramUsernames string `json:"telegram_usernames" form:"telegram_usernames" gorm:"column:telegram_usernames;"`
	TelegramBotToken  string `json:"telegram_bot_token" form:"telegram_bot_token" gorm:"column:telegram_bot_token;not null;"`

	FirstName string `json:"first_name" form:"first_name" gorm:"column:first_name;"`
	LastName  string `json:"last_name" form:"last_name" gorm:"column:last_name;"`
	IsForum   bool   `json:"is_forum" form:"is_forum" gorm:"column:is_forum;"`
	IsBot     bool   `json:"is_bot" form:"is_bot" gorm:"column:is_bot;"`
	IsPremium bool   `json:"is_premium" form:"is_premium" gorm:"column:is_premium;"`

	// 核验请求id
	ById string `json:"by_id" form:"by_id" gorm:"-"`

	Page   int    `json:"-" form:"page" gorm:"-"`
	Size   int    `json:"-" form:"size" gorm:"-"`
	Filter string `json:"-" form:"filter" gorm:"-"`
}
