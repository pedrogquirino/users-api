package processors

import (
	"users-api/internal/core/domain"
	"users-api/internal/core/domain/users"
	"users-api/internal/core/ports"

	"golang.org/x/net/context"
)

type SaveUserProcessor struct {
	repository ports.UsersRepository
}

func NewSaveUserProcessor(options *domain.Options) *SaveUserProcessor {
	return &SaveUserProcessor{
		repository: options.UsersRepository,
	}
}

func (p *SaveUserProcessor) ProcessRequest(ctx context.Context, request users.UserRequest) error {
	user := users.ConvertUserRequestToUser(request)
	user.Password, _ = domain.HashPassword(user.Password)

	err := p.repository.Save(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
