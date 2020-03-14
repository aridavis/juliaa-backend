package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DBConfig struct {
	DBName   string
	Host     string
	User     string
	Port     int
	Password string
}

var DB *gorm.DB

func BuildConfig() *DBConfig {
	return &DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "",
		DBName:   "juliabot",
	}
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
