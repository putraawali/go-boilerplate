package usecase

import (
	"context"
	"service-notif/model"
	"service-notif/user/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepositoryContract
}

func NewUserUsecase(userRepository repository.UserRepositoryContract) UserUsecaseContract {
	return &UserUsecase{
		userRepo: userRepository,
	}
}

func (usecase *UserUsecase) GetByID(ctx context.Context, user_id int) (model.UserModel, error) {
	return usecase.userRepo.GetByID(ctx, user_id)
}
