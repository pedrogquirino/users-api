package processors

import (
	"users-api/internal/core/domain"
	"users-api/internal/core/domain/address"
	"users-api/internal/core/ports"

	"golang.org/x/net/context"
)

type GetAddressProcessor struct {
	repository ports.AddressRepository
}

func NewGetAddressProcessor(options *domain.Options) *GetAddressProcessor {
	return &GetAddressProcessor{
		repository: options.AddressRepository,
	}
}

func (p *GetAddressProcessor) ProcessRequest(ctx context.Context, addressId string, userId string) (*address.AddressResponse, error) {
	user, err := p.repository.GetById(ctx, addressId, userId)
	if err != nil {
		return nil, err
	}
	usersResponse := p.buildResponse(user)
	return usersResponse, nil
}

func (p *GetAddressProcessor) buildResponse(addressData *address.Address) *address.AddressResponse {
	addressResponse := address.AddressResponse{
		Id:       addressData.Id,
		District: addressData.District,
		Street:   addressData.Street,
		Number:   addressData.Number,
		ZipCode:  addressData.ZipCode,
	}
	return &addressResponse
}
