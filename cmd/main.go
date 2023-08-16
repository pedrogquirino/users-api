package main

import (
	"users-api/infra/config"
	"users-api/infra/repository"
	"users-api/infra/server"
	"users-api/internal/core/domain"
	"users-api/internal/core/services/processors"
)

func main() {

	conf := config.New()
	validator := domain.NewValidator()
	options := &domain.Options{
		Config:    conf,
		Validator: validator,
	}
	repository.Start(options)
	processors := processors.NewProcessors(options)
	server.Start(options, processors)
}
