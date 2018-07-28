package model

import (
	"time"

	"github.com/dkeng/cogo/module/store"
)

// BaseModel is base entity
type BaseModel struct {
	ID        int64      `json:"id" gorm:"primary_key;unique_index"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;type:DATETIME"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null;type:DATETIME"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index" gorm:"type:DATETIME"`
}

// AutoMigrate 自动迁移
func AutoMigrate() {
	store.DB.AutoMigrate(
		&Application{},
		&Config{},
	)
}
