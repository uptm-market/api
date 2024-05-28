package fb

import (
	"context"

	v16 "github.com/justwatch/facebook-marketing-api-golang-sdk/marketing/v16"
	"go.mod/entity"
	"go.mod/rest"
)

const (
	appID       = "SUA_APP_ID"
	appSecret   = "bea9e01a01182211c02a7754a2e95b6e"
	accessToken = "EAAF6ELj49ZAABOxyZBWNyIZBZBsNUHMlRwlDnqmKR7tlNVi1dcPtZCMcph3LhCbdX4JTAZCOTos69evof6o8ZB2arCWuvuAl3fbk0xQgli1CXuZCYwX4lRfI7iktjXOcPBNK8LZBag1JrZCRjSKU1zLLVJnAbMxkP3eOSjfju5bL0gbhvk5VwdqiNVNTXhrRYoccmg"
	adAccountID = "ID_DA_CONTA_DE_ANUNCIOS"
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
