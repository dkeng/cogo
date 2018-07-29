package config

import "errors"

var (
	errConfigAdd    = errors.New("配置文件添加出错")
	errConfigExist  = errors.New("该配置项已存在")
	errConfigUpdate = errors.New("修改配置文件出错")
)
