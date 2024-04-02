package entity

import v16 "github.com/justwatch/facebook-marketing-api-golang-sdk/marketing/v16"

type FacebookCampaignAdAccount struct {
	ID                int    `json:"id"`
	CampaignAccountID string `json:"campaign_account_id"`
	AdAccountID       string `json:"ad_account_id"`
	UserID            int    `json:"user_id"`
}

type Campaign struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Location    string      `json:"location"`
	FoundedYear int         `json:"founded_year"`
	Employees   int         `json:"employees"`
	Revenue     int         `json:"revenue"`
	Website     string      `json:"website"`
	Contact     ContactInfo `json:"contact"`
	Tags        []string    `json:"tags"`
}

// ContactInfo represents contact information
type ContactInfo struct {
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type CampaignReturnCount struct {
	Count    int            `json:"count"`
	Campaign []v16.Campaign `json:"campaign"`
}
