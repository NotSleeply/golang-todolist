package main

import (
	"log"
	"todo/config"
	"todo/db"
	"todo/routes"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	db.Init()

	// 设置路由
	r := routes.SetupRoutes()

	// 启动服务器
	log.Printf("服务器启动在端口 %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
