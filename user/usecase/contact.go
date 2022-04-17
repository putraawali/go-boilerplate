package usecase

import (
	"context"
	"service-notif/model"
)

type UserUsecaseContract interface {
	GetByID(ctx context.Context, user_id int) (model.UserModel, error)
}
