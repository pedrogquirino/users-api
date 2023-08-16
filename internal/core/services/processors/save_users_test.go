package processors

import (
	"context"
	"errors"
	"testing"
	"users-api/internal/core/domain"
	"users-api/internal/core/domain/users"
	"users-api/internal/core/ports"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SaveUserProcessorSuite struct {
	suite.Suite
	options        *domain.Options
	processor      *SaveUserProcessor
	userRepository *ports.MockUsersRepository
}

func (s *SaveUserProcessorSuite) initTest() {
	s.userRepository = new(ports.MockUsersRepository)
	s.options = &domain.Options{
		UsersRepository: s.userRepository,
	}
	s.processor = NewSaveUserProcessor(s.options)
}

func TestSaveUserProcessorSuite(t *testing.T) {
	suite.Run(t, new(SaveUserProcessorSuite))
}

func (s *SaveUserProcessorSuite) TestGetUsersProcessorLogic() {
	s.T().Run("TestSaveUserProcessorSuite_SaveUser_Success", func(t *testing.T) {
		s.initTest()
		var (
			user = users.UserRequest{
				Name:  "Pedro",
				Age:   39,
				Email: "pedro@example.com",
			}
			ctx = context.Background()
		)
		s.userRepository.On("Save", ctx, mock.Anything).Return(nil)
		error := s.processor.ProcessRequest(ctx, user)
		assert.Nil(t, error)
	})

	s.T().Run("TestSaveUserProcessorSuite_SaveUser_Error", func(t *testing.T) {
		s.initTest()
		var (
			userRequest = users.UserRequest{
				Name:  "Pedro",
				Age:   39,
				Email: "pedro@example.com",
			}
			ctx = context.Background()
		)
		s.userRepository.On("Save", ctx, mock.Anything).Return(errors.New("error"))
		error := s.processor.ProcessRequest(ctx, userRequest)
		assert.NotNil(t, error)
	})
}
