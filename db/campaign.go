package db

import (
	"context"

	infradb "go.mod/connect"
	"go.mod/entity"
)

func CreateFacebookCampaign(ctx context.Context, data entity.FacebookCampaignAdAccount) error {
	_, err := infradb.Load().ExecContext(ctx, `INSERT INTO facebook_campaign_ad_account (campaign_account_id, ad_account_id, user_id) VALUES($1, $2, $3)`, data.CampaignAccountID, data.CampaignAccountID, data.UserID)
	if err != nil {
		return err
	}
	return nil
}

func ReturnCampaign(ctx context.Context, userId int) ([]entity.FacebookCampaignAdAccount, error) {
	var array []entity.FacebookCampaignAdAccount
	rows, err := infradb.Load().QueryContext(ctx, `SELECT id, campaign_account_id, ad_account_id, user_id FROM facebook_campaign_ad_account where user_id=$1;`, userId)
	if err != nil {
		return nil, err
	}
	var data entity.FacebookCampaignAdAccount
	for rows.Next() {
		rows.Scan(&data.ID, &data.CampaignAccountID, &data.AdAccountID, &data.UserID)
		array = append(array, data)
	}

	return array, nil

}
