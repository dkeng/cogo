package entity

// Application 应用
type Application struct {
	AppKey    string `gorm:"not null;unique_index"`
	AppSecret string `gorm:"not null;size:32"`
	BaseEntity
}
