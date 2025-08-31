package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		DSN string
	}
	Log struct {
		Level      string
		Filename   string
		MaxSize    int
		MaxBackups int
		MaxAge     int
		Compress   bool
	}
	CORS struct {
		AllowedOrigins []string
		AllowedMethods []string
	}
	Redis struct {
		Host string
	}
	Contract struct {
		HttpUrl      string
		WebsocketUrl string
	}
}

var Cfg Config

func LoadConfig() {
	env := os.Getenv("APP_ENV") // 从环境变量读取运行环境
	if env == "" {
		env = "dev" // 默认开发环境
	}

	v := viper.New()
	v.SetConfigName(fmt.Sprintf("config.%s", env))
	v.SetConfigType("yaml")
	v.AddConfigPath("config/")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("读取配置失败: %v", err)
	}
	if err := v.Unmarshal(&Cfg); err != nil {
		log.Fatalf("解析配置失败: %v", err)
	}

	fmt.Printf("✅ 当前运行环境: %s\n", env)
}
