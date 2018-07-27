package store

import (
	"github.com/dkeng/cogo/store/entity"
	"github.com/jinzhu/gorm"
)

// ConfigStore 配置文件存储
type ConfigStore struct {
	baseStore
}

// Init 初始化
func (c *ConfigStore) Init(db *gorm.DB) *ConfigStore {
	c.Db = db
	c.Name = "ConfigStore"
	return c
}

// Insert 插入
func (c *ConfigStore) Insert(config *entity.Config) error {
	return c.Db.Create(config).Error
}

// DeleteByID 根据ID删除配置
func (c *ConfigStore) DeleteByID(id int64) error {
	return c.Db.Where("id = ?", id).Delete(&entity.Config{}).Error
}

// Update 修改
func (c *ConfigStore) Update(config *entity.Config) error {
	return c.Db.Update(config).Error
}
