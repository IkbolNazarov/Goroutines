package repository

import (
	"channels/internal/db"
	"channels/internal/models"
	"fmt"
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
	var results []models.All
	batchSize:=20
	c := make(chan []models.All)
	result  := db.DataB.Table("name").
		Select("name.first_name, name.last_name, info.address, info.phone_numb, info.email, pic.pic").
		Joins("JOIN info ON name.id = info.id").
		Joins("JOIN pic ON name.id = pic.id").
		Where("name.id BETWEEN ? AND ?", begin, end).FindInBatches(&results, batchSize, func(tx *gorm.DB, batch int) error {
			for _, result := range results{
				fmt.Println(tx.RowsAffected)
				fmt.Println(batch)
				c<- results
				go r.ExportToXLS(batchSize, c)
			}
			return nil 
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) ExportToXLS(total int, c chan []models.All) error {
	f := excelize.NewFile()
	all := <-c
	close(c)
	sheetName := "Sheet1"
	columns := []string{"id", "first_name", "last_name", "address", "phone_numb", "email", "pic"}
	for i, colName := range columns {
		cell := fmt.Sprintf("%s%d", string(rune('A'+i)), 1)
		f.SetCellValue(sheetName, cell, colName)
	}
	row := 2
	for i := 1; i <= total; i++ {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), (all)[row-2].Id)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), (all)[row-2].FirstName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), (all)[row-2].LastName)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), (all)[row-2].Address)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), (all)[row-2].PhoneNumb)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), (all)[row-2].Email)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), (all)[row-2].Pic)
		row++
	}

	randomise := strconv.Itoa(rand.Intn(99999))
	err := f.SaveAs(randomise + "output.xlsx")
	if err != nil {
		return err
	}
	return nil
}

