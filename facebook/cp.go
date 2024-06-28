package fb

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Cp(token string) string {

	// URL da API com a variável do token
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/me?fields=adaccounts{campaigns{id,name}}&access_token=%s", token)

	// Crie uma nova solicitação HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erro ao criar a solicitação:", err)
		return ""
	}

	// Crie um cliente HTTP
	client := &http.Client{}

	// Envie a solicitação
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar a solicitação:", err)
		return ""
	}
	defer resp.Body.Close()

	// Leia a resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return ""
	}

	// Retorne a resposta como string
	return string(body)
}
