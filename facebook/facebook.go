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

func InitConfig(ctx context.Context) (*v16.Service, error) {
	var fbService *v16.Service
	data, err := db.ReturnCampaign(ctx, int(ctx.Value("userid").(uint)))
	if err != nil {
		rest.LogError(err, "Erro ao criar conexao com api do facebook, problema ao consultar db")
		return nil, err
	}

	tk, err := db.ReturnTokenFacebook(ctx, uint(data.UserID))
	if err != nil {
		rest.LogError(err, "Erro ao criar conexao com api do facebook, problema ao consultar db")
		return nil, err
	}

	log.Println(tk, "token")
	log.Println(data.AppSecret, "app_secret")

	if data != nil {
		fbService, err = v16.New(nil, tk, appSecret)
		if err != nil {
			rest.LogError(err, "Erro ao criar conexao com api do facebook")
			return nil, context.TODO().Err()
		}
	}

	return fbService, nil
}

func Init(ctx context.Context, userId int) *v16.Service {
	tk, err := db.ReturnTokenFacebook(ctx, uint(userId))
	if err != nil {
		rest.LogError(err, "Erro ao criar conexao com api do facebook, problema ao consultar db")
		return nil
	}

	fbService, err := v16.New(nil, tk, appSecret)
	if err != nil {
		rest.LogError(err, "Erro ao criar conexao com api do facebook")
		return nil
	}

	return fbService
}
