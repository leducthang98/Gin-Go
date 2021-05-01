package controller

import (
	"github.com/go-chi/chi"
	"go-graphql-boilderplate/pkg/api_helper"
	usecase "go-graphql-boilderplate/usecase/user"
	"net/http"
)

func MountUserRouters(r *chi.Mux, userService usecase.IUser) {
	userController := UserController{
		userService: userService,
	}
	r.Route("/user", func(r chi.Router) {
		r.Get("/all", userController.getAllUser)

	})
}

type UserController struct {
	userService usecase.IUser
}

func (u *UserController) getAllUser(w http.ResponseWriter, r *http.Request) {
	if users, err := u.userService.GetAll(); err != nil {
		api_helper.WriteResponse(w, http.StatusBadRequest, nil)
	} else {
		api_helper.WriteResponse(w, http.StatusOK, users)
	}

}
