package main

import (
	"fmt"
	"github.com/1255177148/golangTask4/config"
	"github.com/1255177148/golangTask4/internal/bootstrap"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"go.uber.org/zap"
)

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
	bootstrap.InitRedis()                                        // 初始化redis

	err := r.Run(":" + config.Cfg.Server.Port) // 启动服务
	if err != nil {
		panic(err)
	}
}
