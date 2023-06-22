package repository

import (
	"channels/internal/db"
	"channels/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{connection: conn}
}


func (r *Repository) GetRecords(c chan []models.All, begin int, end int) (error) {
	var result []models.All 
	defer close(c)

	err := db.DataB.Model(&models.Name{}).
    Select("name.first_name, name.last_name, info.address, info.phone_numb, info.email, pic.pic").
    Joins("JOIN info ON name.id = info.id").
    Joins("JOIN pic ON name.id = pic.id").
    Where("name.id BETWEEN ? AND ?", begin, end).
    Find(&result).Error


if err != nil {
    return err
}
	c <- result
	return nil
}