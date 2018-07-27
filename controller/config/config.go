package config

import (
	"github.com/dkeng/cogo/model"
	"github.com/dkeng/cogo/module/store"
	"github.com/gin-gonic/gin"
)

// Get 获取
func Get(*gin.Context) {

}

// Post 新增
func Post(*gin.Context) {

}

// Delete 删除
func Delete(*gin.Context) {
	store.DB.Where("id = ?", "").Delete(&model.Config{})
}

// Put 修改
func Put(*gin.Context) {
	conf := new(model.Config)
	store.DB.Update(conf)
}
