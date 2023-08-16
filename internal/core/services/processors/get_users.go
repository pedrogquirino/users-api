package processors

import (
	"users-api/internal/core/domain"
	"users-api/internal/core/domain/users"
	"users-api/internal/core/ports"

	"golang.org/x/net/context"
)

type GetUsersProcessor struct {
	repository ports.UsersRepository
}

func NewGetUsersProcessor(options *domain.Options) *GetUsersProcessor {
	return &GetUsersProcessor{
		repository: options.UsersRepository,
	}
}

func (p *GetUsersProcessor) ProcessRequest(ctx context.Context, userId string) (*users.UserResponse, error) {
	user, err := p.repository.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}
	usersResponse := p.buildResponse(user)
	return usersResponse, nil
}

func (p *GetUsersProcessor) buildResponse(user *users.User) *users.UserResponse {
	usersResponse := users.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Age:      user.Age,
		Email:    user.Email,
		Address:  user.Address,
		Password: user.Password,
	}
	return &usersResponse
}
