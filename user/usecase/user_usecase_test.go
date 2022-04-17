package usecase

import (
	"context"
	"service-notif/model"
	"service-notif/user/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = repository.UserRepoMock{Mock: mock.Mock{}}
var userUsecase = NewUserUsecase(&userRepository)

func TestSuccessFindById(t *testing.T) {
	var (
		ctx  = context.Background()
		id   = 1
		user = model.User{
			Id:            1,
			TransporterId: 1,
			ShipperId:     1,
		}
	)

	userRepository.Mock.On("FindById", ctx, id).Return(user, nil)
	user, err := userUsecase.FindById(ctx, id)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestNotFound(t *testing.T) {
	var (
		ctx  = context.Background()
		id   = 1
		user = model.User{
			Id:            0,
			TransporterId: 0,
			ShipperId:     0,
		}
	)

	userRepository.Mock.On("FindById", ctx, id).Return(user, nil)
	user, err := userUsecase.FindById(ctx, id)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}
