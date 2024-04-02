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

func CampaignRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/campaign", middleware.AuthMiddlewareWithClaims(http.HandlerFunc(createCampaignHandler)))
	r.Get("/campaign", middleware.AuthMiddlewareWithClaims(http.HandlerFunc(returnCampaignHandler)))

	return r
}

func createCampaignHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.URL.Query().Get("user_id")
	type Body struct {
		CampaignAccountID string `json:"campaign_account_id"`
		AdAccountID       string `json:"ad_account_id"`
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
		UserID:            id,
		CampaignAccountID: body.CampaignAccountID,
		AdAccountID:       body.AdAccountID,
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
	manager := core.NewUserCampaign()
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	send, err := manager.Get(ctx, userIdInt)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, send)
}
