package store

import "github.com/jinzhu/gorm"

// ApplicationStore 应用存储
type ApplicationStore struct {
	baseStore
}

// Init 初始化
func (a *ApplicationStore) Init(db *gorm.DB) *ApplicationStore {
	a.Db = db
	a.Name = "ApplicationStore"
	return a
}
