package fb

import (
	"context"
	"log"

	v16 "github.com/justwatch/facebook-marketing-api-golang-sdk/marketing/v16"
	"go.mod/db"
	"go.mod/rest"
)

const (
	appID       = "SUA_APP_ID"
	appSecret   = "121ad7ed5049737df0336cfc20e26f72"
	accessToken = "EAAfXJXDJoCkBO1mqglUJFVRawEZC1YPZBOvXbXRiLZC1QDUgE0gxLgU4ZBYoz4J8Qvsyns3pPrslkeZBisdl63YPQ6JOdfQ9yQ0fuzrvz8X3oLaWi7IGXZBDqjONsEntsjD23ueNmpKR0hua1yuVOzgODWuGguKYvC1UgKwQhGMrSkO0LoJzzdSsoLyUWv1rTakGDo5gzAnvQbOKn4"
	adAccountID = "ID_DA_CONTA_DE_ANUNCIOS"
	campaignID  = "ID_DA_CAMPANHA"
)

func InitConfig(ctx context.Context) *v16.Service {
	var fbService *v16.Service
	data, err := db.ReturnCampaign(ctx, int(ctx.Value("userid").(uint)))
	if err != nil {
		rest.LogError(err, "Erro ao criar conexao com api do facebook, problema ao consultar db")
		return nil
	}

	tk, err := db.ReturnTokenFacebook(ctx, uint(data.UserID))
	if err != nil {
		rest.LogError(err, "Erro ao criar conexao com api do facebook, problema ao consultar db")
		return nil
	}

	log.Println(tk, "token")
	log.Println(data.AppSecret, "app_secret")

	if data != nil {
		fbService, err = v16.New(nil, tk, appSecret)
		if err != nil {
			rest.LogError(err, "Erro ao criar conexao com api do facebook")
			return nil
		}
	}

	return fbService
}

// func Init(ctx context.Context, ID entity.FacebookCampaignAdAccount) []v16.Campaign {

// 	fbService, err := v16.New(nil, accessToken, appSecret)
// 	if err != nil {
// 		rest.LogError(err, "Erro ao criar conexao com api do facebook")
// 		return nil
// 	}

// 	id := ID.Token

// 	campaigns, err := fbService.Campaigns.List(id).Do(ctx)
// 	if err != nil {
// 		rest.LogError(err, "Erro ao retornar dados da campanha")
// 		return nil
// 	}
// 	return campaigns
// }
