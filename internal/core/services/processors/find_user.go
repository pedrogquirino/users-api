package processors

import (
	"users-api/internal/core/domain"
	"users-api/internal/core/domain/users"
	"users-api/internal/core/ports"

	"golang.org/x/net/context"
)

type FindUserProcessor struct {
	repository ports.UsersRepository
}

func NewFindUserProcessor(options *domain.Options) *FindUserProcessor {
	return &FindUserProcessor{
		repository: options.UsersRepository,
	}
}

func (p *FindUserProcessor) ProcessRequest(ctx context.Context) (*[]users.UserResponse, error) {
	user, err := p.repository.Find(ctx)
	if err != nil {
		return nil, err
	}
	usersResponse := p.buildResponse(user)
	return usersResponse, nil
}

func (p *FindUserProcessor) buildResponse(usersData []users.User) *[]users.UserResponse {
	var usersResponse []users.UserResponse
	for _, user := range usersData {
		usersResponse = append(usersResponse, users.UserResponse{
			Id:       user.Id,
			Name:     user.Name,
			Age:      user.Age,
			Email:    user.Email,
			Address:  user.Address,
			Password: user.Password,
		})
	}
	return &usersResponse
}
