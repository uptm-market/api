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
	// r.Put("/", middleware.AuthMiddleware(getIndexHandlerFunc(updateCampaign)))

	return r
}

func createCampaignHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.URL.Query().Get("user_id")
	type Body struct {
		CampaignAccountID string   `json:"app_secret"`
		AdAccountID       string   `json:"token"`
		BusinessID        []string `json:"businessID"`
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
		BusinessID: body.BusinessID,
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
	cpID := r.URL.Query().Get("cp_id")
	manager := core.NewUserCampaign()
	if cpID == "" {
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			rest.SendError(w, err)
			return
		}
		send, err := manager.List(ctx, userIdInt)
		if err != nil {
			rest.SendError(w, err)
			return
		}
		rest.Send(w, send)
	} else {
		data, err := manager.Get(ctx, cpID)
		if err != nil {
			rest.SendError(w, err)
			return
		}
		rest.Send(w, data)
	}

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
