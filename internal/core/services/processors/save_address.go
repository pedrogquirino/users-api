package processors

import (
	"users-api/internal/core/domain"
	"users-api/internal/core/domain/address"
	"users-api/internal/core/ports"

	"golang.org/x/net/context"
)

type SaveAddressProcessor struct {
	repository ports.AddressRepository
}

func NewSaveAddressProcessor(options *domain.Options) *SaveAddressProcessor {
	return &SaveAddressProcessor{
		repository: options.AddressRepository,
	}
}

func (p *SaveAddressProcessor) ProcessRequest(ctx context.Context, userId string, request address.AddressRequest) error {
	address := address.ConvertAddressRequestToAddress(request)

	err := p.repository.Save(ctx, userId, address)
	if err != nil {
		return err
	}
	return nil
}
