package processors

import (
	"users-api/internal/core/domain"
	"users-api/internal/core/domain/address"
	"users-api/internal/core/domain/users"
)

type Processors struct {
	FindUserProcessor    users.FindUserProcessor
	GetUsersProcessor    users.GetUsersProcessor
	SaveUserProcessor    users.SaveUserProcessor
	FindAddressProcessor address.FindAddressProcessor
	GetAddressProcessor  address.GetAddressProcessor
	SaveAddressProcessor address.SaveAddressProcessor
}

func NewProcessors(options *domain.Options) *Processors {
	return &Processors{
		FindUserProcessor:    NewFindUserProcessor(options),
		GetUsersProcessor:    NewGetUsersProcessor(options),
		SaveUserProcessor:    NewSaveUserProcessor(options),
		FindAddressProcessor: NewFindAddressProcessor(options),
		GetAddressProcessor:  NewGetAddressProcessor(options),
		SaveAddressProcessor: NewSaveAddressProcessor(options),
	}
}
