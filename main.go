package main

import (
	"github.com/debuginn/GoShortLink/dao"
	"github.com/debuginn/GoShortLink/model"
	"github.com/debuginn/GoShortLink/router"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// 导入配置文件
	initConfig()
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
	port := viper.GetString("server.port")
	panic(r.Run(":" + port))
}

// 导入系统配置文件
func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
