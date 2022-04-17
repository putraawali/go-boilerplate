package usecase

import (
	"context"
	"service-notif/model"
)

type UserUsecaseContract interface {
	FindById(ctx context.Context, user_id int) (model.User, error)
}
