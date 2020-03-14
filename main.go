package main

import (
	"Telegraf/Config"
	Models "Telegraf/Models/Reminder"
	"Telegraf/Routes"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildConfig()))
	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Reminder{})

	r := Routes.SetupRouter()
	r.Run()

}
