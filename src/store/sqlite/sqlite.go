package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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

func switchDB(tran *gorm.DB, db *gorm.DB) *gorm.DB {
	if tran != nil {
		return tran
	}
	if db != nil {
		return db
	}
	panic(errors.New("转换数据库失败"))
}
