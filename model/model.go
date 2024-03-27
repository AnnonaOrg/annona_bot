package model

// import (
// 	"time"
// 	// "gorm.io/gorm"
// )

type Model struct {
	// ID uint64 `gorm:"primarykey;autoIncrement" json:"id"`
	// CreatedAt time.Time      `json:"-"`
	// UpdatedAt time.Time      `json:"-"`
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type BaseModel struct {
	// Id uint64 `gorm:"column:id" json:"id"`
	// CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	// UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
	// // DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}

// Token represents a json web token
type Token struct {
	// Token string `json:"token"`
}
