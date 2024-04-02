package entity

type FacebookCampaignAdAccount struct {
	ID                int    `json:"id"`
	CampaignAccountID string `json:"campaign_account_id"`
	AdAccountID       string `json:"ad_account_id"`
	UserID            int    `json:"user_id"`
}
