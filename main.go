package main

import (
	"GoShortLink/dao"
	"GoShortLink/model"
	"GoShortLink/route"
)

func main() {
	// 初始化数据库
	err := dao.InitDB()
	if err != nil {
		panic(err)
	}
	// 服务结束关闭数据库
	defer dao.CloseDB()
	// 模型绑定
	dao.DB.AutoMigrate(&model.User{})
	// 注册路由
	r := routes.SetupRouter()
	// 开启服务
	panic(r.Run())
}
