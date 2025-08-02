package main

import (
	"fmt"
	"github.com/1255177148/golangTask4/config"
	"github.com/1255177148/golangTask4/docs"
	"github.com/1255177148/golangTask4/internal/bootstrap"
	"github.com/1255177148/golangTask4/internal/utils/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title           个人博客系统后端 API 文档
// @version         1.0
// @description     这是一个基于 Gin 的 个人博客系统后端API
// @host            localhost:9080
// @BasePath        /api
func main() {
	config.LoadConfig() // 加载配置文件

	fmt.Println("✅ 配置加载成功")
	fmt.Println("数据库连接:", config.Cfg.Database.DSN)

	bootstrap.InitLogger()
	log.Init(bootstrap.Logger)
	defer func(Logger *zap.Logger) {
		err := Logger.Sync()
		if err != nil {
			panic(err)
		}
	}(bootstrap.Logger)

	bootstrap.InitDB()                                           // 初始化 GORM和sqlx
	r := bootstrap.InitApp(bootstrap.DB.Gorm, bootstrap.DB.Sqlx) // 初始化 Gin
	// 注册 Swagger
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	bootstrap.InitRedis() // 初始化redis

	err := r.Run(":" + config.Cfg.Server.Port) // 启动服务
	if err != nil {
		panic(err)
	}
}
