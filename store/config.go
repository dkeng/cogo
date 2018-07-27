package store

import "github.com/jinzhu/gorm"

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
