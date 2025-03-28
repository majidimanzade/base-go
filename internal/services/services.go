package services

import "github.com/majidimanzade/base-go/internal/services/generator"

type Services struct {
	Generator generator.GeneratorService
}

func NewServices() *Services {
	return &Services{
		Generator: generator.New(),
	}
}
