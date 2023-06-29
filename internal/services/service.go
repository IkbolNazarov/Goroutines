package services

import (
	"channels/internal/repository"
	"log"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}

func (s *Services) GetUser(begin int, end int) error {
	go s.Repository.GetRecords(begin, end)

	log.Println("Done")
	return nil
}
