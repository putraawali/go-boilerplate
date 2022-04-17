package httphandler

import (
	"context"
	"net/http"
	"service-notif/model"
	"service-notif/user/usecase"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userUsecase usecase.UserUsecaseContract
}

func NewUserHandler(userUsecase usecase.UserUsecaseContract) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (handler *UserHandler) GetRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/{id}", handler.GetByID)

	return router
}

func (handler *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	//! For url parameters
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		model.ResponseNOK(w, http.StatusBadRequest, model.ErrBadParamInput.Error())
		return
	}
	//! For query parameters
	// id := r.URL.Query().Get("id")
	ctx := context.Background()
	data, err := handler.userUsecase.GetByID(ctx, id)

	if err != nil {
		model.ResponseNOK(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}
	if data.Id == 0 {
		model.ResponseNOK(w, http.StatusNotFound, model.ErrNotFound.Error())
		return
	}

	model.ResponseOK(w, http.StatusOK, data)
}
