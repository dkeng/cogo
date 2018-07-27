package config

import (
	"github.com/dkeng/cogo/model"
	"github.com/gin-gonic/gin"
)

// Post 新增
func Post(*gin.Context) {

}

// Delete 删除
func Delete(*gin.Context) {
	model.DB.Where("id = ?", "").Delete(&model.Config{})
}

// Put 修改
func Put(*gin.Context) {
	conf := new(model.Config)
	model.DB.Update(conf)
}
