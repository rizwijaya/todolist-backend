package database

import (
	"fmt"
	"log"
	"todolist-backend/infrastructures/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQL() *gorm.DB {
	config := config.New()
	//Connection to Databases
	dsn := config.Database.User + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.Dbname + "?charset=utf8&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Connection to Databases Success")
	//Migration Databases
	err = Db.AutoMigrate(&Activities{}, &Todos{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Migration Databases Success")

	return Db
}

func NewDatabase() *gorm.DB {
	return MySQL()
}
