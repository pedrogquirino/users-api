package address

import (
	"context"
)

type FindAddressProcessor interface {
	ProcessRequest(ctx context.Context, userId string) (*[]AddressResponse, error)
}

type GetAddressProcessor interface {
	ProcessRequest(ctx context.Context, userId string, id string) (*AddressResponse, error)
}

type SaveAddressProcessor interface {
	ProcessRequest(ctx context.Context, userId string, address AddressRequest) error
}
