package generator

import (
	"log"
)

type GeneratorService interface {
	Run() error
}
type Service struct{}

func New() *Service {
	return &Service{}
}

func (service *Service) Run() error {
	log.Println("Running generator service")
	return nil
}
