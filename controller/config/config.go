package config

import (
	"github.com/dkeng/cogo/model"
	"github.com/dkeng/cogo/module/store"
	"github.com/gin-gonic/gin"
)

// Get 获取
func Get(c *gin.Context) {

}

// Post 新增
func Post(c *gin.Context) {

}

// Delete 删除
func Delete(c *gin.Context) {
	store.DB.Where("id = ?", "").Delete(&model.Config{})
}

// Put 修改
func Put(c *gin.Context) {
	conf := new(model.Config)
	store.DB.Update(conf)
}
