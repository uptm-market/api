package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	v16 "github.com/justwatch/facebook-marketing-api-golang-sdk/marketing/v16"
	"go.mod/core"
	"go.mod/entity"
	"go.mod/middleware"
	"go.mod/rest"
)

func CampaignRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", middleware.AuthMiddleware(getIndexHandlerFunc(returnCampaignHandler)))
	r.Post("/", middleware.AuthMiddleware(getIndexHandlerFunc(createCampaignHandler)))
	r.Post("/copy", middleware.AuthMiddleware(getIndexHandlerFunc(cloneCampaignHandler)))
	r.Put("/active/{id}", middleware.AuthMiddleware(getIndexHandlerFunc(activeHandler)))
	r.Get("/list/businessid/{userId}", middleware.AuthMiddleware(getIndexHandlerFunc(listBusinessHandler)))
	r.Get("/listAll/{userId}", middleware.AuthMiddleware(getIndexHandlerFunc(listBusinessAll)))
	// r.Put("/", middleware.AuthMiddleware(getIndexHandlerFunc(updateCampaign)))

	return r
}

func createCampaignHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.URL.Query().Get("user_id")
	type Body struct {
		CampaignAccountID string            `json:"app_secret"`
		AdAccountID       string            `json:"token"`
		Business          []entity.Business `json:"business"`
	}
	var body Body
	if err := rest.ParseBody(w, r, &body); err != nil {
		return
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		return
	}
	bodyData := entity.FacebookCampaignAdAccount{
		UserID:     id,
		AppSecret:  body.CampaignAccountID,
		Token:      &body.AdAccountID,
		BusinessID: body.Business,
	}

	manager := core.NewUserCampaign()
	err = manager.Create(ctx, bodyData)
	if err != nil {
		rest.SendError(w, err)
		return
	}

	rest.Send(w, nil)

}

func returnCampaignHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.URL.Query().Get("user_id")
	// cpID := r.URL.Query().Get("cp_id")
	manager := core.NewUserCampaign()

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	send, err := manager.ListAds(ctx, uint(userIdInt))
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, send)

}

func cloneCampaignHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var data v16.Campaign
	if err := rest.ParseBody(w, r, &data); err != nil {
		rest.SendError(w, err)
		return
	}
	manager := core.NewUserCampaign()
	err := manager.CreateCampaignFull(ctx, data)
	if err != nil {
		rest.SendError(w, err)
		return
	}
}

// func updateCampaign(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	var data v16.Ad
// 	if err := rest.ParseBody(w, r, &data); err != nil {
// 		rest.SendError(w, err)
// 		return
// 	}
// 	manager := core.NewUserCampaign()
// 	// err := manager.updateCampaign(ctx, data)
// }

func activeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	// idint, err := strconv.ParseInt(id, 10, 10)
	// if err != nil {
	// 	rest.SendError(w, err)
	// 	return
	// }
	manager := core.NewUserCampaign()
	err := manager.Active(ctx, id)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, http.StatusOK)

}

func listBusinessHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := chi.URLParam(r, "userId")
	idint, err := strconv.ParseInt(userId, 10, 10)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	manager := core.NewUserCampaign()
	ListBusiness, err := manager.ListBusinessId(ctx, int(idint))
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, ListBusiness)
}

func listBusinessAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := chi.URLParam(r, "userId")
	idint, err := strconv.ParseInt(userId, 10, 10)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	manager := core.NewUserCampaign()
	send, err := manager.GetAllBusiness(ctx, int(idint))
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, send)
}
