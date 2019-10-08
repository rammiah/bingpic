package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("mysql",
		"bing:bing@tcp(rammiah.org:3306)/bing?parseTime=True&charset=utf8mb4&timeout=3s")
	if err != nil {
		log.Errorf("connect to db failed, error: %v", err)
		panic(err)
	}
	//log.Printf("connect to db success")
	log.Infof("connect to db success")
}
