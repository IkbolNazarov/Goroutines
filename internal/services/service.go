package services

import (
	"channels/internal/models"
	"channels/internal/repository"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}

func (s *Services) GetUser(begin int, end int) (*[]models.All, error) {
	var wg sync.WaitGroup
	all := make([]models.All, end-begin+1)
	for i := begin; i <= end; i++ {
		wg.Add(1)
		go func(i int) (*[]models.All, error) {
			res, err := s.Repository.GetUser(i)
			if err != nil {
				return nil, err
			}
			all[i-begin] = *res
			wg.Done()
			return &all, nil
		}(i)
	}
	wg.Wait()
	return &all, nil
}

func (s *Services) ExportToXLS(all *[]models.All) error {
	f := excelize.NewFile()
	sheetName := "Sheet1"

	columns := []string{"id", "first_name", "last_name", "address", "phone_numb", "email", "pic"}
	for i, colName := range columns {
		cell := fmt.Sprintf("%s%d", string('A'+i), 1)
		f.SetCellValue(sheetName, cell, colName)
	}
	total := len(*all)
	log.Println(total)
	row := 2
	for i := row; i <= total+1; i++ {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), (*all)[i-2].Id)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), (*all)[i-2].FirstName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), (*all)[i-2].LastName)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), (*all)[i-2].Address)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), (*all)[i-2].PhoneNumb)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), (*all)[i-2].Email)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), (*all)[i-2].Pic)
		row++
	}
	randomise := strconv.Itoa(rand.Intn(99999))
	err := f.SaveAs(randomise + "output.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
