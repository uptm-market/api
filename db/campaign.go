package db

import (
	"context"
	"database/sql"

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
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO facebook_campaign_ad_account_token (token_id, user_id)
VALUES ($1, $2)
ON CONFLICT (user_id) 
DO UPDATE SET token_id = EXCLUDED.token_id;
`, data.Token, data.UserID)

	if err != nil {
		return err
	}
	return nil
}

func ReturnTokenFacebook(ctx context.Context, userId uint) (token string, err error) {
	err = infradb.Get().QueryRowContext(ctx, `select token_id from facebook_campaign_ad_account_token where user_id=$1`, userId).Scan(&token)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return token, err
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
