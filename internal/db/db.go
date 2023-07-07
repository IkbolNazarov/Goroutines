package db

import (
	"channels/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DataB *gorm.DB

func GetDbConnection()		 *gorm.DB {
	_, dbConf, err := config.GetConfig("db")
	conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConf.DbConnection.Host, dbConf.DbConnection.Port, dbConf.DbConnection.User, dbConf.DbConnection.Password, dbConf.DbConnection.Dbname)
	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("не удалось подключиться к базе данных")
	}
	DataB = db
	return db
}
