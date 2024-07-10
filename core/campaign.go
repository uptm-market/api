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

func (c *UserCampaign) CreateCampaignFull(ctx context.Context, data v16.Campaign) error {
	arrayReturnMain, err := fb.InitConfig(ctx)
	if err != nil {
		return err
	}
	_, err = arrayReturnMain.Campaigns.Create(ctx, data)
	if err != nil {
		return rest.LogError(err, "c.UserCampaign.CreateCampaignfull fb.Create")
	}

	return nil
}

func (c *UserCampaign) List(ctx context.Context, userId int) (map[string]interface{}, error) {
	// var array []string
	// var arrayReturnCam *v16.CampaignListCall
	// ar, err := db.ReturnCampaign(ctx, userId)
	// if err != nil {
	// 	return nil, rest.LogError(err, "ReturnCampaign")
	// }
	// log.Println("001")
	// for _, a := range ar.BusinessID {
	// 	log.Println("teste entrou", a)
	// 	arrayReturnMain, err := fb.InitConfig(ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	arrayReturn, err := arrayReturnMain.AdAccounts.List(ctx, a)
	// 	if err != nil {
	// 		return nil, &rest.Error{Status: 400, Code: "bad_request_fb_lib", Message: err.Error()}
	// 	}
	// 	log.Println("001 - meio")
	// 	log.Println("teste abah", arrayReturn)
	// 	log.Println("-------f--------")
	// 	// Check if arrayReturn has enough elements before accessing index i

	// 	for _, r := range arrayReturn {
	// 		log.Println("001 - meio - loop - array")
	// 		array = append(array, r.AccountID)
	// 	}

	// }

	// log.Println("002")

	// log.Println(array)
	// arrayReturnMain, err := fb.InitConfig(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// arrayReturnCam = arrayReturnMain.Campaigns.List("680041130165805")
	// log.Println("002 - meio")

	// log.Println("003")

	tk, err := db.ReturnTokenFacebook(ctx, uint(userId))
	if err != nil {
		rest.LogError(err, "Erro ao criar conexao com api do facebook, problema ao consultar db")
		return nil, err
	}
	st := fb.Cp(tk)
	return st, nil

}

func (c *UserCampaign) Get(ctx context.Context, campaign string) (*v16.Campaign, error) {
	arrayReturnMain, err := fb.InitConfig(ctx)
	if err != nil {
		return nil, err
	}
	data, err := arrayReturnMain.Campaigns.Get(ctx, campaign)
	if err != nil {
		return nil, &rest.Error{Status: 400, Code: "bad_request_fb_lib", Message: err.Error()}
	}

	return data, nil
}

// func (c *UserCampaign) updateCampaign(ctx context.Context, data v16.Ad) {
// 	err := fb.InitConfig().Ads.Update(ctx, da)
// }

func (c *UserCampaign) Active(ctx context.Context, id int) (err error) {
	err = db.Active(ctx, id)
	if err != nil {
		return
	}
	return nil
}

func (c *UserCampaign) ListBusinessId(ctx context.Context, id int) (*entity.FacebookCampaignAdAccount, error) {
	data, err := db.ReturnCampaign(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *UserCampaign) GetAllBusiness(ctx context.Context, id int) ([]entity.Business, error) {
	data, err := db.ListBusinessHandler(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
