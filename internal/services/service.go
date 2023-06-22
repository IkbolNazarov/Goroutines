package services

import (
	"channels/internal/models"
	"channels/internal/repository"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}


func (s *Services) GetUser() error {
	c := make(chan []models.All)
	go s.Repository.GetRecords(c)
	all, ok := <-c
	if ok {
		s.ExportToXLS(len(all), &all)
	}
	return nil
}

func (s *Services) ExportToXLS(total int, all *[]models.All) error {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	columns := []string{"id", "first_name", "last_name", "address", "phone_numb", "email", "pic"}
	for i, colName := range columns {
		cell := fmt.Sprintf("%s%d", string(rune('A'+i)), 1)
		f.SetCellValue(sheetName, cell, colName)
	}
	row := 2
	for i := 1; i <= total; i++ {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), (*all)[row-2].Id)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), (*all)[row-2].FirstName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), (*all)[row-2].LastName)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), (*all)[row-2].Address)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), (*all)[row-2].PhoneNumb)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), (*all)[row-2].Email)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), (*all)[row-2].Pic)
		row++
	}
	randomise := strconv.Itoa(rand.Intn(99999))
	err := f.SaveAs(randomise + "output.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
