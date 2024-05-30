package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go.mod/core"
	"go.mod/entity"
	"go.mod/middleware"
	"go.mod/rest"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", userinfoHandler)
	r.Get("/me", middleware.AuthMiddleware(getIndexHandlerFunc(userinfoMeHandler)))
	r.Post("/", crateUserHandler)
	r.Post("/login", loginHandler)
	r.Put("/{id}", middleware.AuthMiddleware(getIndexHandlerFunc(updateUserHandler)))
	r.Put("/{id}/password", middleware.AuthMiddleware(getIndexHandlerFunc(updatePasswordHandler)))
	return r
}

func crateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user entity.UserCreations
	if err := rest.ParseBody(w, r, &user); err != nil {
		return
	}
	manager := core.NewUserManager()
	err := manager.CreateUser(ctx, user)
	if err != nil {
		rest.SendError(w, err)
		return
	}
}

func userinfoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	send, err := core.NewUserManager().User(ctx, id)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, send)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var login entity.Login
	if err := rest.ParseBody(w, r, &login); err != nil {
		return
	}
	manager := core.NewUserManager()
	token, err := manager.Login(ctx, login.Email, login.Password)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, token)
}

func userinfoMeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	manager := core.NewUserManager()
	data, err := manager.GetMeInfo(ctx)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, data)

}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var data entity.UserUpdated
	if err := rest.ParseBody(w, r, &data); err != nil {
		return
	}
	id := chi.URLParam(r, "id")
	manager := core.NewUserManager()
	ids, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	err = manager.UpdatedUser(ctx, data, uint(ids))
	if err != nil {
		rest.SendError(w, err)
		return
	}

}

func updatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type Updated struct {
		OldPassowrd string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}

	var data Updated
	if err := rest.ParseBody(w, r, &data); err != nil {
		return
	}
	json := entity.UpdatePassword{
		OldPassword: data.OldPassowrd,
		NewPassword: data.NewPassword,
	}
	id := chi.URLParam(r, "id")
	manager := core.NewUserManager()
	ids, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	err = manager.UpdatedPassowrd(ctx, json, uint(ids))
	if err != nil {
		rest.SendError(w, err)
		return
	}

}
