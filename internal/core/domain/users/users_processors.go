package users

import (
	"context"
)

type FindUserProcessor interface {
	ProcessRequest(ctx context.Context) (*[]UserResponse, error)
}

type GetUsersProcessor interface {
	ProcessRequest(ctx context.Context, userId string) (*UserResponse, error)
}

type SaveUserProcessor interface {
	ProcessRequest(ctx context.Context, user UserRequest) error
}
