package entity

// Config 配置
type Config struct {
	Key   string `gorm:"not null;unique_index"`
	Value string `gorm:"not null"`
	// 模式
	Mode byte `gorm:"not null;type:char"`
	// 版本
	Version float64 `gorm:"not null"`
	// 应用ID
	AppID int64 `gorm:"not null"`
	BaseEntity
}
