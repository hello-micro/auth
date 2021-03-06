package database

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/hello-micro/auth/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// Conn gorm.DB
	Conn *gorm.DB
	// once sync.Once
	once sync.Once
)

// Init init db
func Init() {
	once.Do(func() {
		var err error
		var dsn string
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
			config.GetDefaultEnv("DB_USERNAME", "root"),
			config.GetDefaultEnv("DB_PASSWORD", ""),
			config.GetDefaultEnv("DB_HOST", "127.0.0.1"),
			config.GetDefaultEnv("DB_PORT", "3306"),
			config.GetDefaultEnv("DB_DATABASE", "root"),
			strings.Replace(config.GetDefaultEnv("APP_TIMEZONE", "Local"), "/", "%2F", -1),
		)
		//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		Conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		//Conn, err = gorm.Open("mysql", dsn)
		if err != nil {
			log.Panicf("init DB err:%v,dsn:%v", err, dsn)
		}
		Conn.Debug()
		// 全局禁用表名复数
		//Conn.SingularTable(true)
		//启用Logger，显示详细日志，默认情况下会打印发生的错误
		//Conn.LogMode(false)
		//Conn.DB().SetMaxIdleConns(10)
		//Conn.DB().SetMaxOpenConns(100)
	})
}
