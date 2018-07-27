package store

import (
	"github.com/dkeng/cogo/store/entity"
	"github.com/jinzhu/gorm"
)

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

// Insert 插入
func (a *ApplicationStore) Insert(app *entity.Application) error {
	return a.Db.Create(app).Error
}

// DeleteByID 删除
func (a *ApplicationStore) DeleteByID(id int64) error {
	return a.Db.Where("id = ?", id).Delete(&entity.Application{}).Error
}
