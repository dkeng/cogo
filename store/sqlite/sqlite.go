package sqlite

import (
	"github.com/jinzhu/gorm"
)

// AllSqliteStore mysql存储
type AllSqliteStore struct {
}

// Init 初始化
func (m *AllSqliteStore) Init(db *gorm.DB) *AllSqliteStore {

	return m
}

// baseSqliteStore shared DB data
type baseSqliteStore struct {
	Db   *gorm.DB
	Name string
}

func getLimitOffset(page, perPage *int) *int {
	if *page <= 0 {
		*page = 1
	}
	if *perPage <= 0 {
		*perPage = 10
	}
	offset := (*page - 1) * *perPage
	return &offset
}
