package application

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/dkeng/pkg/convert"

	"github.com/dkeng/pkg/context/gin"

	"github.com/dkeng/cogo/model"
	"github.com/dkeng/cogo/module/store"
)

// Get 获取
func Get(w *gin.WrapContenxt) {
	var apps []model.Application
	if store.DB.Find(apps).RecordNotFound() {
		w.Status(http.StatusNotFound)
	} else {
		w.OKJSON(gin.Result{
			"data":  apps,
			"total": len(apps),
		})
	}
}

// GetOne 获取一个
func GetOne(w *gin.WrapContenxt) {
	var config model.Config
	if store.DB.First(&config, w.Param("id")).RecordNotFound() {
		w.Status(http.StatusNotFound)
	} else {
		w.OKJSON(config)
	}
}

// GetConfigs 获取配置文件
func GetConfigs(w *gin.WrapContenxt) {
	appID := w.Param("id")
	name, nameExist := w.GetQuery("name")
	query := "app_id = ?"
	values := []interface{}{
		appID,
	}
	if nameExist {
		query += " and name = ?"
		values = append(values, name)
	}

	var configs []model.Config
	if err := store.DB.Where(query, values...).Find(&configs).Error; err != nil {
		w.ErrorJSON(errAppSelect)
		return
	}
	w.OKJSON(gin.Result{
		"data":  configs,
		"total": len(configs),
	})
}

// Post 创建
func Post(w *gin.WrapContenxt) {
	app := new(model.Application)
	h := md5.New()
	h.Write([]byte(convert.ToString(time.Now().Unix())))
	app.AppSecret = hex.EncodeToString(h.Sum(nil))

	if err := store.DB.Create(app).Error; err != nil {
		log.Println(err)
		w.ErrorJSON(errAppCreate)
	} else {
		w.OKJSON(app)
	}
}

// Delete 删除
func Delete(w *gin.WrapContenxt) {
	if err := store.DB.Where("id = ?", "").Delete(&model.Application{}).Error; err != nil {
		log.Println(err)
		w.ErrorJSON(errAppDelete)
	} else {
		w.Status(http.StatusNoContent)
	}
}
