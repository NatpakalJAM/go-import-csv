package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // driver
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Db is Database Instance
var Db *gorm.DB

// Init connect database
func Init() {
	var err error

	// `ragnarok_rtc_s2` connect
	confRtc := viper.GetStringMapString("db")
	connRtc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		confRtc["username"],
		confRtc["password"],
		confRtc["host"],
		confRtc["port"],
		confRtc["database"])

	Db, err = gorm.Open("mysql", connRtc)
	if viper.GetString("Environment") != "production" {
		Db.LogMode(true)
	}
	if err != nil {
		panic(fmt.Errorf("Error connect db `` %s", err))
	}
}
