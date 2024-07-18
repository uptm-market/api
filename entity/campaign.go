package entity

import v16 "github.com/justwatch/facebook-marketing-api-golang-sdk/marketing/v16"

type FacebookCampaignAdAccount struct {
	ID         int        `json:"id"`
	AppSecret  string     `json:"app_secret"`
	Token      *string    `json:"token"`
	UserID     int        `json:"user_id"`
	BusinessID []Business `json:"business_id"`
	Act        string     `json:"act"`
}
type Business struct {
	Name string `json:"name"`
	ID   string `json:"id"`
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

type AdAccount struct {
	AccountID string `json:"account_id"`
	ID        string `json:"id"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type Paging struct {
	Cursors Cursors `json:"cursors"`
}

type OwnedAdAccounts struct {
	Data   []AdAccount `json:"data"`
	Paging Paging      `json:"paging"`
}

type Response struct {
	OwnedAdAccounts OwnedAdAccounts `json:"owned_ad_accounts"`
	ID              string          `json:"id"`
}
