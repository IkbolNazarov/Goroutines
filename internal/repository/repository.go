package repository

import (
	"channels/internal/db"
	"channels/internal/models"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/gorm"
)

type Repository struct {
	connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{connection: conn}
}

// func (r *Repository) GetRecords(c chan []models.All, begin int, end int) (error) {
// 	var result []models.All
// 	defer close(c)
// 	err := db.DataB.Model(&models.Name{}).
// Select("name.first_name, name.last_name, info.address, info.phone_numb, info.email, pic.pic").
// Joins("JOIN info ON name.id = info.id").
// Joins("JOIN pic ON name.id = pic.id").
// Where("name.id BETWEEN ? AND ?", begin, end).
//     Find(&result).Limit(100).Error
// if err != nil {
//     return err
// }
// 	c <- result
// 	return nil
// }

func (r *Repository) GetRecords(begin int, end int) error {
	f := excelize.NewFile()
	batchSize := 100
	c := make(chan []models.Name)
	go func() {
		defer close(c)
		for offset := begin; offset <= end; offset += batchSize {
			var results []models.Name
			result := db.DataB.Offset(offset).Limit(batchSize).Find(&results)
			if result.Error != nil {
				fmt.Println(result.Error)
				return
			}
			c <- results
		}
	}()
	r.ExportToXLS(end-begin+1, c, f)
	r.SaveXLS(f)
	return nil
}

func (r *Repository) ExportToXLS(total int, c chan []models.Name, f *excelize.File) {
	sheetName := "Sheet1"
	columns := []string{"id", "first_name", "last_name"}
	for i, colName := range columns {
		cell := fmt.Sprintf("%s%d", string(rune('A'+i)), 1)
		f.SetCellValue(sheetName, cell, colName)
	}
	row := 2
	for results := range c {
		for _, result := range results {
			f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), result.Id)
			f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), result.FirstName)
			f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), result.LastName)
			row++
		}
	}
}

func (r *Repository) SaveXLS(f *excelize.File) error {
	randomize := strconv.Itoa(rand.Intn(99999))
	err := f.SaveAs(randomize + "output.xlsx")
	log.Println("output ok")
	if err != nil {
		fmt.Println("Ошибка при сохранении файла:", err)
		return err
	}
	return nil
}
