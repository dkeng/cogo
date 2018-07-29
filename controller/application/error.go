package application

import "errors"

var (
	errAppCreate = errors.New("创建应用失败")
	errAppDelete = errors.New("删除应用失败")
	errAppSelect = errors.New("查询应用失败")
)
