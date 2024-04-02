package db

import (
	"context"

	infradb "go.mod/connect"
	"go.mod/entity"
)

func CreateFacebookCampaign(ctx context.Context, data entity.FacebookCampaignAdAccount) error {
	_, err := infradb.DB.ExecContext(ctx, `INSERT INTO facebook_campaign_ad_account (campaign_account_id, ad_account_id, user_id) VALUES($1, $2, $3)`, data.CampaignAccountID, data.CampaignAccountID, data.UserID)
	if err != nil {
		return err
	}
	return nil
}
