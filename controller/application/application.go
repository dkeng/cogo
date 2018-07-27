package application

import (
	"github.com/dkeng/cogo/model"
	"github.com/gin-gonic/gin"
)

// Post 创建
func Post(*gin.Context) {
	app := new(model.Application)
	model.DB.Create(app)
}

// Delete 删除
func Delete(*gin.Context) {
	model.DB.Where("id = ?", "").Delete(&model.Application{})
}
