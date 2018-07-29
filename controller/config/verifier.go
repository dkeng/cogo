package config

import (
	"github.com/dkeng/cogo/model"
	"github.com/dkeng/cogo/module/store"
)

type postConfModel struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Value string `json:"value" form:"value" binding:"required"`
	// 模式
	Mode byte `json:"mode" form:"mode" binding:"required"`
	// 版本
	Version float64 `json:"version" form:"version" binding:"required"`
	// 应用ID
	AppID int64 `json:"app_id" form:"app_id" binding:"required"`
}

func (c *postConfModel) Validation(obj interface{}) error {
	var count int64
	if err := store.DB.Model(&model.Config{}).Where(c).Count(&count).Error; err != nil {
		return errConfigAdd
	}
	if count > 0 {
		return errConfigExist
	}
	return nil
}

type putConfModel struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Value string `json:"value" form:"value" binding:"required"`
	// 模式
	Mode byte `json:"mode" form:"mode" binding:"required"`
	// 版本
	Version float64 `json:"version" form:"version" binding:"required"`
	// 应用ID
	AppID int64 `json:"app_id" form:"app_id" binding:"required"`
}

func (c *putConfModel) Validation(obj interface{}) error {
	var count int64
	if err := store.DB.Model(&model.Config{}).Where(c).Where("id != ?", obj).Count(&count).Error; err != nil {
		return errConfigAdd
	}
	if count > 0 {
		return errConfigExist
	}
	return nil
}
