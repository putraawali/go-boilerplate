package user

import (
	"database/sql"
	user_handler "service-notif/user/http_handler"
	"service-notif/user/repository"
	"service-notif/user/usecase"
)

func NewUserInjection(db *sql.DB) *user_handler.UserHandler {
	userRepository := repository.UserConnection(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := user_handler.NewUserHandler(userUsecase)

	return userHandler
}
