package fb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mod/entity"
)

func Cp(token, act string) map[string]interface{} {
	// URL da API com a variável do token
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s?fields=adaccounts{campaigns{id,name,status,account_id,budget_rebalance_flag,buying_type,created_time,lifetime_budget,issues_info,source_campaign,special_ad_category,special_ad_category_country,start_time,stop_time,daily_budget,budget_remaining}}&access_token=%s", act, token)

	// Função para fazer a solicitação e ler a resposta
	fetchData := func(url string) (map[string]interface{}, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar a solicitação: %v", err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("erro ao enviar a solicitação: %v", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler a resposta: %v", err)
		}

		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter bytes para JSON: %v", err)
		}

		return result, nil
	}

	// Função para extrair o link de paginação "next"
	getNextURL := func(result map[string]interface{}) (string, bool) {
		if adAccounts, ok := result["adaccounts"].(map[string]interface{}); ok {
			if paging, ok := adAccounts["paging"].(map[string]interface{}); ok {
				if next, ok := paging["next"].(string); ok {
					return next, true
				}
			}
		}
		return "", false
	}

	// Variável para armazenar todos os resultados
	allResults := make(map[string]interface{})
	allResults["adaccounts"] = []interface{}{}

	// Loop de paginação
	for {
		result, err := fetchData(url)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// Adicione os dados da resposta atual aos resultados acumulados
		if adAccounts, ok := result["adaccounts"].(map[string]interface{}); ok {
			if data, ok := adAccounts["data"].([]interface{}); ok {
				allResults["adaccounts"] = append(allResults["adaccounts"].([]interface{}), data...)
			}
		}

		// Verifique se há mais páginas para buscar
		nextURL, hasNext := getNextURL(result)
		if !hasNext {
			break
		}
		url = nextURL
	}

	return allResults
}

func CpByBusinessID(token string, businessId string) *entity.OwnedAdAccounts {
	business := businessId
	accessToken := token
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s?fields=owned_ad_accounts&access_token=%s", business, accessToken)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response status", resp.Status)
		return nil
	}

	var result *entity.OwnedAdAccounts
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	return result
}
