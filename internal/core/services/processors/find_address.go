package processors

import (
	"users-api/internal/core/domain"
	"users-api/internal/core/domain/address"
	"users-api/internal/core/ports"

	"golang.org/x/net/context"
)

type FindAddressProcessor struct {
	repository ports.AddressRepository
}

func NewFindAddressProcessor(options *domain.Options) *FindAddressProcessor {
	return &FindAddressProcessor{
		repository: options.AddressRepository,
	}
}

func (p *FindAddressProcessor) ProcessRequest(ctx context.Context, userId string) (*[]address.AddressResponse, error) {
	address, err := p.repository.Find(ctx, userId)
	if err != nil {
		return nil, err
	}
	addressResponse := p.buildResponse(address)
	return addressResponse, nil
}

func (p *FindAddressProcessor) buildResponse(addressData []address.Address) *[]address.AddressResponse {
	var addressResponse []address.AddressResponse
	for _, addressItem := range addressData {
		addressResponse = append(addressResponse, address.AddressResponse{
			Id:       addressItem.Id,
			District: addressItem.District,
			Street:   addressItem.Street,
			Number:   addressItem.Number,
			ZipCode:  addressItem.ZipCode,
		})
	}
	return &addressResponse
}
