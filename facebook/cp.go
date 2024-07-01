package fb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Cp(token string) map[string]interface{} {

	// URL da API com a variável do token
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/me?fields=adaccounts{campaigns{id,name}}&access_token=%s", token)

	// Crie uma nova solicitação HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erro ao criar a solicitação:", err)
		return nil
	}

	// Crie um cliente HTTP
	client := &http.Client{}

	// Envie a solicitação
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar a solicitação:", err)
		return nil
	}
	defer resp.Body.Close()

	// Leia a resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return nil
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Erro ao converter bytes para JSON:", err)
	} else {
		fmt.Println("Resultado da conversão de bytes para JSON:", result)
	}
	// Retorne a resposta como string
	return result
}
