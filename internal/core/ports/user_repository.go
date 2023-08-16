package ports

import (
	"context"
	"users-api/internal/core/domain/users"
)

type UsersRepository interface {
	Find(ctx context.Context) ([]users.User, error)
	GetById(ctx context.Context, id string) (*users.User, error)
	Save(ctx context.Context, user *users.User) error
}
