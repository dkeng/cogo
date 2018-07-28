package model

// Config 配置
type Config struct {
	Name  string `json:"name" gorm:"not null;unique_index"`
	Value string `json:"value" gorm:"not null"`
	// 模式
	Mode byte `json:"mode" gorm:"not null;type:char"`
	// 版本
	Version float64 `json:"version" gorm:"not null"`
	// 应用ID
	AppID int64 `json:"app_id" gorm:"not null"`
	BaseModel
}
