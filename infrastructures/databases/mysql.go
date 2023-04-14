package database

import (
	"log"
	"todolist-backend/infrastructures/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQL() *gorm.DB {
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	dsn := config.Database.User + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.Dbname + "?charset=utf8&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}
	return Db
}

func NewDatabase() *gorm.DB {
	return MySQL()
}
