package model

import (
	"time"
)

// BaseModel is base entity
type BaseModel struct {
	ID        int64      `gorm:"primary_key;unique_index"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;type:DATETIME"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null;type:DATETIME"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index" gorm:"type:DATETIME"`
}
