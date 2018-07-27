package application

import (
	"github.com/dkeng/cogo/model"
	"github.com/dkeng/cogo/module/store"
	"github.com/gin-gonic/gin"
)

// Get 获取
func Get(*gin.Context) {

}

// Post 创建
func Post(*gin.Context) {
	app := new(model.Application)
	store.DB.Create(app)
}

// Delete 删除
func Delete(*gin.Context) {
	store.DB.Where("id = ?", "").Delete(&model.Application{})
}
