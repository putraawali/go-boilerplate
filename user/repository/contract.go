package repository

import (
	"context"
	"service-notif/model"
)

type UserRepositoryContract interface {
	GetByID(ctx context.Context, user_id int) (model.UserModel, error)
}
