package ports

import (
	"context"
	"users-api/internal/core/domain/address"
)

type AddressRepository interface {
	Find(ctx context.Context, userId string) ([]address.Address, error)
	GetById(ctx context.Context, id string, userId string) (*address.Address, error)
	Save(ctx context.Context, userId string, address *address.Address) error
}
