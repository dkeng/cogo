package model

// Application 应用
type Application struct {
	AppSecret string `json:"secret" gorm:"column:secret;not null;size:32"`
	BaseModel
}
