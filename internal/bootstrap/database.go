// 初始化数据库

package bootstrap

import (
	"github.com/1255177148/golangTask4/config"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBInstance struct {
	Gorm *gorm.DB
	Sqlx *sqlx.DB
}

var DB DBInstance

// InitDB 初始化数据库
func InitDB() {
	dsn := config.Cfg.Database.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewZapGormLogger(Logger), // ✅ 使用自定义 zap GORM logger
	})
	if err != nil {
		//log.Fatalf("GORM连接失败: %v", err)
		log.Error("GORM 连接失败", zap.Error(err))
	}
	sqlxDB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Error("SQLX 连接失败", zap.Error(err))
	}
	if err = sqlxDB.Ping(); err != nil {
		log.Error("SQLX Ping 失败", zap.Error(err))
	}
	DB = DBInstance{Gorm: db, Sqlx: sqlxDB}
	log.Info("✅ 数据库初始化完成")
}
