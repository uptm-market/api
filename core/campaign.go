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

func (c *UserCampaign) Get(ctx context.Context, userId int) ([]v16.Campaign, error) {
	var arrayReturn []v16.Campaign
	ar, err := db.ReturnCampaign(ctx, userId)
	if err != nil {
		return nil, rest.LogError(err, "ReturnCampaign")
	}
	size := len(ar)
	for i := 0; i < size; i++ {
		returnBody := fb.Init(ctx, ar[i])
		arrayReturn = append(arrayReturn, returnBody[i])
	}

	return arrayReturn, nil

}
