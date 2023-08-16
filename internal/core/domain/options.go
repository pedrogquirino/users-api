package domain

import (
	"users-api/internal/core/ports"

	"users-api/infra/config"
)

type Options struct {
	Config            *config.Config
	Validator         *Validator
	UsersRepository   ports.UsersRepository
	AddressRepository ports.AddressRepository
}
