package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/1255177148/golangTask4/config"
	"github.com/1255177148/golangTask4/docs"
	"github.com/1255177148/golangTask4/internal/bootstrap"
	"github.com/1255177148/golangTask4/internal/container"
	"github.com/1255177148/golangTask4/internal/utils/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title           个人博客系统后端 API 文档
// @version         1.0
// @description     这是一个基于 Gin 的 个人博客系统后端API
// @host            localhost:9080
// @BasePath        /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 1. 加载配置
	config.LoadConfig() // 加载配置文件
	fmt.Println("✅ 配置加载成功")
	fmt.Println("数据库连接:", config.Cfg.Database.DSN)

	// 2. 初始化日志
	bootstrap.InitLogger()
	log.Init(bootstrap.Logger)
	defer func(Logger *zap.Logger) {
		err := Logger.Sync()
		if err != nil {
			panic(err)
		}
	}(bootstrap.Logger)

	// 3. 初始化ethereum client
	bootstrap.InitContractClient() // 初始化ethereum client
	// 4. 初始化数据库
	bootstrap.InitDB() // 初始化 GORM和sqlx
	// 5. 初始化 Gin
	r := bootstrap.InitApp(bootstrap.DB.Gorm, bootstrap.DB.Sqlx) // 初始化 Gin
	// 6. 注册 Swagger
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 7. 初始化 Redis
	bootstrap.InitRedis() // 初始化redis

	// 8. 使用 http.Server 包装 Gin，支持优雅关闭
	srv := &http.Server{
		Addr:    ":" + config.Cfg.Server.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("listen: " + err.Error())
		}
	}()
	fmt.Println("🚀 服务已启动，端口:", config.Cfg.Server.Port)
	// 9. 监听系统信号 (Ctrl+C, kill)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("⚠️ 收到退出信号，开始优雅关闭...")
	// 10. 关闭链上listener
	container.StopAllListeners()
	// 11. 优雅关闭 Gin
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("❌ Gin 优雅关闭失败:", err)
	} else {
		fmt.Println("✅ Gin 已优雅退出")
	}
}
