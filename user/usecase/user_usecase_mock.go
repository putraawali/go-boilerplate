package usecase

import (
	"context"
	"service-notif/model"

	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (m *UserUsecaseMock) FindById(ctx context.Context, id int) (model.User, error) {
	args := m.Called(id)

	return args.Get(0).(model.User), args.Error(1)
}
