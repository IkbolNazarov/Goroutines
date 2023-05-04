package repository

import (
	"channels/internal/db"
	"channels/internal/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{connection: conn}
}


func (r *Repository) GetUserByChan(id int, c chan models.All) (chan models.All, error) {
	log.Println(id)
	var all models.All
	tx := db.DataB.Table("name").Where("id = ?", id).Find(&all)
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println("name: ", all.FirstName)

	tx = db.DataB.Table("pic").Where("id = ?", id).Find(&all)
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println("pic: ", all.Pic)
	tx = db.DataB.Table("info").Where("id = ?", id).Find(&all)
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println("info: ", all.Email)
	c<-all
	return c, nil
}
