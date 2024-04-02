package core

import (
	"context"

	"go.mod/db"
	"go.mod/entity"
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
