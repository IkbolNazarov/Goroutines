package services

import (
	"channels/internal/repository"
	"errors"
	"log"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}

func (s *Services) GetUser(begin int, end int) error {
	err:= s.Repository.GetRecords(begin, end)
	if err!= nil {
		return errors.New("Error in repository")
	}
	log.Println("Done")
	return nil
}