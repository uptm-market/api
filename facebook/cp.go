package fb

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Cp(token, act string) map[string]interface{} {
	// URL da API com a variável do token
	// token = "EAAfXJXDJoCkBO9kJHIyIS10LNKLqvtxLVZAKYMKJNZCKuLizltSpEf8Jf1glZBJfxlHyqEbAa2tPZAWIpymoRlfEDltgFaCZCxKLCij8LfBjD9XcZAUZAOzZAblnXViLgLQwdt0ysB796erMXUEdC9ABOEhUyauQvsrUtZB4vzZBNVkAhGeOQyjiZCqAzCknqo0vVtXsDpeh1PuTFrw6ZBSVYEQ1qAwqwMuyKX3ZBkksZD"
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s?fields=ads{campaign,bid_amount,effective_status,creative,configured_status},amount_spent,business_city,min_daily_budget&access_token=%s", act, token)
	log.Println(url)
	// Função para fazer a solicitação e ler a resposta

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("erro ao criar a solicitação: %v", err)
		return nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("erro ao criar a solicitação: %v", err)
		return nil
	}

	// defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("erro ao ler a resposta: %v", err)
		return nil
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Errorf("erro ao converter bytes para JSON: %v", err)
		return nil
	}

	return result

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

func CpByBusinessID(token string, businessID string) []string {
	// token = "EAAfXJXDJoCkBO9kJHIyIS10LNKLqvtxLVZAKYMKJNZCKuLizltSpEf8Jf1glZBJfxlHyqEbAa2tPZAWIpymoRlfEDltgFaCZCxKLCij8LfBjD9XcZAUZAOzZAblnXViLgLQwdt0ysB796erMXUEdC9ABOEhUyauQvsrUtZB4vzZBNVkAhGeOQyjiZCqAzCknqo0vVtXsDpeh1PuTFrw6ZBSVYEQ1qAwqwMuyKX3ZBkksZD"
	// businessID = "7042491049135964"
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s?fields=owned_ad_accounts&access_token=%s", businessID, token)
	fmt.Println("URL:", url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Erro ao criar solicitação:", err)
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Erro ao fazer solicitação:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Erro: Recebido status diferente de 200 OK:", resp.Status)
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Erro ao ler o corpo da resposta:", err)
		return nil
	}

	var response Response
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		return nil
	}
	log.Println(response)
	var strarray []string
	for _, account := range response.OwnedAdAccounts.Data {
		strarray = append(strarray, account.ID)
	}
	log.Println(strarray)
	return strarray
}
