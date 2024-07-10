package db

import (
	"context"
	"database/sql"
	"log"

	infradb "go.mod/connect"
	"go.mod/entity"
)

func CreateFacebookCampaign(ctx context.Context, data entity.FacebookCampaignAdAccount) error {
	for _, business := range data.BusinessID {
		_, err := infradb.Get().ExecContext(ctx, `INSERT INTO facebook_campaign_ad_account (app_secret, business_id, user_id) VALUES($1, $2, $3)`, data.AppSecret, business, data.UserID)
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
func ReturnCampaign(ctx context.Context, userId int) (*entity.FacebookCampaignAdAccount, error) {
	// Define the query to fetch the main data
	queryMain := `SELECT MIN(id) AS min_id, MIN(app_secret) AS min_app_secret, MIN(user_id) AS min_user_id
FROM facebook_campaign_ad_account
WHERE user_id = $1  and active =true ;
`
	// Define the query to fetch the business IDs
	queryBusiness := `SELECT business_id, name FROM facebook_campaign_ad_account WHERE user_id=$1 and active =true`

	// Initialize a variable to hold the main data
	var data entity.FacebookCampaignAdAccount

	// Fetch the main data
	err := infradb.Get().QueryRowContext(ctx, queryMain, userId).Scan(&data.ID, &data.AppSecret, &data.UserID)
	if err != nil {
		return nil, err
	}

	// Fetch the business IDs
	rows, err := infradb.Get().QueryContext(ctx, queryBusiness, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Collect the business IDs into a slice
	var strArray []entity.Business
	for rows.Next() {
		var bid entity.Business
		if err := rows.Scan(&bid.ID, &bid.Name); err != nil {
			return nil, err
		}
		strArray = append(strArray, bid)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Assign the business IDs to the data
	data.BusinessID = strArray
	log.Println("teste aqui")
	return &data, nil
}

func Active(ctx context.Context, id int) (err error) {
	_, err = infradb.Get().ExecContext(ctx, `UPDATE facebook_campaign_ad_account
SET active = NOT active
WHERE id = $1;
`, id)
	return err
}

func ListBusinessHandler(ctx context.Context, userId int) ([]entity.Business, error) {
	queryBusiness := `SELECT business_id, name FROM facebook_campaign_ad_account WHERE user_id=$1`
	rows, err := infradb.Get().QueryContext(ctx, queryBusiness, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Collect the business IDs into a slice
	var strArray []entity.Business
	for rows.Next() {
		var bid entity.Business
		if err := rows.Scan(&bid.ID, &bid.Name); err != nil {
			return nil, err
		}
		strArray = append(strArray, bid)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return strArray, nil
}
