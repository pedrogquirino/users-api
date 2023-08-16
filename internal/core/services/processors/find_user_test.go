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

type FindUserProcessorSuite struct {
	suite.Suite
	options        *domain.Options
	processor      *FindUserProcessor
	userRepository *ports.MockUsersRepository
}

func (s *FindUserProcessorSuite) initTest() {
	s.userRepository = new(ports.MockUsersRepository)
	s.options = &domain.Options{
		UsersRepository: s.userRepository,
	}
	s.processor = NewFindUserProcessor(s.options)
}

func TestFindUserProcessorSuite(t *testing.T) {
	suite.Run(t, new(FindUserProcessorSuite))
}

func (s *FindUserProcessorSuite) TestFindUserProcessorLogic() {
	s.T().Run("TestFindUserProcessorSuite_FindUser_Success", func(t *testing.T) {
		s.initTest()
		var (
			users = []users.User{{Name: "Pedro", Age: 39}, {Name: "Pedro", Age: 39}}
			ctx   = context.Background()
		)
		s.userRepository.On("Find", ctx).Return(users, nil)
		_, error := s.processor.ProcessRequest(ctx)
		assert.Nil(t, error)
	})

	s.T().Run("TestFindUserProcessorSuite_FindUser_Error", func(t *testing.T) {
		s.initTest()
		var (
			ctx = context.Background()
		)
		s.userRepository.On("Find", ctx).Return(nil, errors.New("error"))
		_, error := s.processor.ProcessRequest(ctx)
		assert.NotNil(t, error)
	})
}
