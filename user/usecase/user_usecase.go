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

func (usecase *UserUsecase) FindById(ctx context.Context, user_id int) (model.User, error) {
	return usecase.userRepo.FindById(ctx, user_id)
}
