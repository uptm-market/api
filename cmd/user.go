package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.mod/core"
	"go.mod/entity"
	"go.mod/middleware"
	"go.mod/rest"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", userinfoHandler)
	r.Get("/me", middleware.AuthMiddlewareWithClaims(http.HandlerFunc(userinfoMeHandler)))
	r.Post("/", crateUserHandler)
	r.Post("/login", loginHandler)
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
		rest.LogError(err, "Error creating user")
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
	rest.Send(w, "token-basic: "+token)
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
