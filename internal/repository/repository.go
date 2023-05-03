package repository

import (
	"channels/internal/db"
	"channels/internal/models"
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{connection: conn}
}

func (r *Repository) GetUser(id int) (*models.All, error) {
	var all models.All
	log.Println(id)
	tx := db.DataB.Table("name").Where("id = ?", id).Find(&all)
	if tx.Error != nil {
		return nil, tx.Error
	}

	
	tx = db.DataB.Table("pic").Where("id = ?", id).Find(&all)
	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = db.DataB.Table("info").Where("id = ?", id).Find(&all)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &all, nil
}
