package fb

import (
	"context"

	v16 "github.com/justwatch/facebook-marketing-api-golang-sdk/marketing/v16"
	"go.mod/entity"
	"go.mod/rest"
)

const (
	appID       = "SUA_APP_ID"
	appSecret   = "408e0b6ac5214145859aa8b7dc44525b"
	accessToken = "ebde8dc251d3a2d3d4642c06c5570b39"
	campaignID  = "ID_DA_CAMPANHA"
)

func InitConfig() *v16.Service {
	fbService, err := v16.New(nil, accessToken, appSecret)
	if err != nil {
		rest.LogError(err, "Erro ao criar conexao com api do facebook")
		return nil
	}

	return fbService
}

func Init(ctx context.Context, ID entity.FacebookCampaignAdAccount) []v16.Campaign {

	fbService, err := v16.New(nil, accessToken, appSecret)
	if err != nil {
		rest.LogError(err, "Erro ao criar conexao com api do facebook")
		return nil
	}

	id := ID.AdAccountID

	campaigns, err := fbService.Campaigns.List(id).Do(ctx)
	if err != nil {
		rest.LogError(err, "Erro ao retornar dados da campanha")
		return nil
	}
	return campaigns
}
