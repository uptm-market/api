package core

import (
	"context"

	v16 "github.com/justwatch/facebook-marketing-api-golang-sdk/marketing/v16"
	"go.mod/db"
	"go.mod/entity"
	fb "go.mod/facebook"
	"go.mod/rest"
)

type UserCampaign struct {
}

func NewUserCampaign() *UserCampaign {
	return &UserCampaign{}
}

func (c *UserCampaign) Create(ctx context.Context, body entity.FacebookCampaignAdAccount) error {
	if err := db.CreateFacebookCampaign(ctx, body); err != nil {
		return rest.LogError(err, "CreateFacebookCampaign")
	}
	return nil
}

func (c *UserCampaign) CreateCampaignFull(ctx context.Context, data v16.Campaign) error {
	_, err := fb.InitConfig().Campaigns.Create(ctx, data)
	if err != nil {
		return rest.LogError(err, "c.UserCampaign.CreateCampaignfull fb.Create")
	}

	return nil
}

func (c *UserCampaign) List(ctx context.Context, userId int) (*v16.CampaignListCall, error) {
	var array []string
	var arrayReturnCam *v16.CampaignListCall
	ar, err := db.ReturnCampaign(ctx, userId)
	if err != nil {
		return nil, rest.LogError(err, "ReturnCampaign")
	}
	for i, a := range ar {
		arrayReturn, err := fb.InitConfig().AdAccounts.List(ctx, a.BusinessID[i])
		if err != nil {
			return nil, &rest.Error{Status: 400, Code: "bad_request_fb_lib", Message: err.Error()}
		}
		array = append(array, arrayReturn[i].AccountID)
	}
	for _, a := range array {
		arrayReturnCam = fb.InitConfig().Campaigns.List(a)

	}
	return arrayReturnCam, nil

}

func (c *UserCampaign) Get(ctx context.Context, campaign string) (*v16.Campaign, error) {
	data, err := fb.InitConfig().Campaigns.Get(ctx, campaign)
	if err != nil {
		return nil, &rest.Error{Status: 400, Code: "bad_request_fb_lib", Message: err.Error()}
	}

	return data, nil
}

// func (c *UserCampaign) updateCampaign(ctx context.Context, data v16.Ad) {
// 	err := fb.InitConfig().Ads.Update(ctx, da)
// }
