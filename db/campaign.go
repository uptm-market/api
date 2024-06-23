package db

import (
	"context"

	infradb "go.mod/connect"
	"go.mod/entity"
)

func CreateFacebookCampaign(ctx context.Context, data entity.FacebookCampaignAdAccount) error {
	for _, business := range data.BusinessID {
		_, err := infradb.Get().ExecContext(ctx, `INSERT INTO facebook_campaign_ad_account (token_id,app_secret, business_id, user_id) VALUES($1, $2, $3, $4)`, data.Token, data.AppSecret, business, data.UserID)
		if err != nil {
			return err
		}

	}
	return nil
}

func ReturnCampaign(ctx context.Context, userId int) ([]entity.FacebookCampaignAdAccount, error) {
	var array []entity.FacebookCampaignAdAccount
	rows, err := infradb.Get().QueryContext(ctx, `SELECT id, token_id,app_secret, business_id, user_id FROM facebook_campaign_ad_account where user_id=$1;`, userId)
	if err != nil {
		return nil, err
	}
	var data entity.FacebookCampaignAdAccount
	for rows.Next() {
		rows.Scan(&data.ID, &data.AppSecret, &data.Token, &data.UserID)
		array = append(array, data)
	}

	return array, nil

}
