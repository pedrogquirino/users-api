package processors

import (
	"context"
	"errors"
	"testing"
	"users-api/internal/core/domain"
	"users-api/internal/core/domain/users"
	"users-api/internal/core/ports"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetUsersProcessorSuite struct {
	suite.Suite
	options        *domain.Options
	processor      *GetUsersProcessor
	userRepository *ports.MockUsersRepository
}

func (s *GetUsersProcessorSuite) initTest() {
	s.userRepository = new(ports.MockUsersRepository)
	s.options = &domain.Options{
		UsersRepository: s.userRepository,
	}
	s.processor = NewGetUsersProcessor(s.options)
}

func TestGetUsersProcessorSuite(t *testing.T) {
	suite.Run(t, new(GetUsersProcessorSuite))
}

func (s *GetUsersProcessorSuite) TestGetUsersProcessorLogic() {
	s.T().Run("TestGetUsersProcessorSuite_GetUser_Success", func(t *testing.T) {
		s.initTest()
		var (
			userId = "0012b0f7-a13c-4e3f-a7c7-42b4522be6f5"
			user   = users.User{Name: "Pedro", Age: 39}
			ctx    = context.Background()
		)
		s.userRepository.On("GetById", ctx, userId).Return(&user, nil)
		_, error := s.processor.ProcessRequest(ctx, userId)
		assert.Nil(t, error)
	})

	s.T().Run("TestGetUsersProcessorSuite_GetUser_Error", func(t *testing.T) {
		s.initTest()
		var (
			userId = "0012b0f7-a13c-4e3f-a7c7-42b4522be6f5"
			ctx    = context.Background()
		)
		s.userRepository.On("GetById", ctx, userId).Return(nil, errors.New("error"))
		_, error := s.processor.ProcessRequest(ctx, userId)
		assert.NotNil(t, error)
	})
}
