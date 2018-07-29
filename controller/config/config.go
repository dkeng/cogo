package config

import (
	"net/http"

	"github.com/dkeng/cogo/model"
	"github.com/dkeng/cogo/module/store"
	"github.com/dkeng/pkg/context/gin"
)

// Get 获取
func Get(w *gin.WrapContenxt) {

}

// Post 新增
func Post(w *gin.WrapContenxt) {
	cm := new(postConfModel)
	if !w.BindValidation(cm, nil) {
		return
	}

	conf := new(model.Config)
	conf.AppID = cm.AppID
	conf.Mode = cm.Mode
	conf.Name = cm.Name
	conf.Value = cm.Value
	conf.Version = cm.Version
	if err := store.DB.Create(conf).Error; err != nil {
		w.ErrorJSON(errConfigAdd)
	} else {
		w.Status(http.StatusCreated)
	}
}

// Delete 删除
func Delete(w *gin.WrapContenxt) {
	store.DB.Where("id = ?", w.Param("id")).Delete(&model.Config{})
}

// Put 修改
func Put(w *gin.WrapContenxt) {
	conf := new(putConfModel)
	if !w.BindValidation(conf, w.Param("id")) {
		return
	}
	if err := store.DB.Model(&model.Config{}).Update(conf).Error; err != nil {
		w.ErrorJSON(errConfigUpdate)
	} else {
		w.Status(http.StatusCreated)
	}
}
