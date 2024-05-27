package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretkey = "SeCuCGdto5Bu7notHb5EvEV0I27pW78PePRLznFG6iV5iHAW2RA9BYDTSj7sotvMuWR63TT1LrvD+yKzyyDKFg=="

type TokenUser struct {
	UserId uint `json:"user"`
}

// CreateToken cria o token de autenticação.
func CreateToken(ctx context.Context, usuarioID uint, email string, role int) (string, error) {
	// Definir claims do token
	claims := jwt.MapClaims{
		"userid":     usuarioID,
		"email":      email,
		"role":       role,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Expira em 24 horas
		"authorized": true,
	}

	// Crie o token com as claims e o método de assinatura HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assine o token com a chave secreta
	tokenString, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(r *http.Request) (bool, error) {
	tokenString := extraction(r)
	if tokenString == "" {
		return false, errors.New("Token de autenticação não fornecido")
	}
	log.Println("Token extraído:", tokenString)

	token, err := jwt.Parse(tokenString, returnKeyVerification)
	if err != nil {
		log.Println("Erro ao analisar o token:", err)
		return false, fmt.Errorf("Erro: %v", err)
	}

	// Verificar manualmente a expiração do token
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expirationTime := int64(claims["exp"].(float64))
		if time.Now().Unix() > expirationTime {
			return false, errors.New("Token de autenticação expirado")
		}
	} else {
		return false, errors.New("Erro ao acessar as reivindicações do token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Exemplo de verificação personalizada nas claims do token
		if authorized, ok := claims["authorized"].(bool); !ok || !authorized {
			return false, errors.New("Token de autenticação não autorizado")
		}
		log.Println("Token válido, claims:", claims)

		// Obter o ID do usuário do token
		userIDFloat, ok := claims["userid"].(float64)
		if !ok {
			return false, errors.New("Erro ao recuperar o ID do usuário do token")
		}
		userID := uint(userIDFloat)

		// Colocar o ID do usuário no contexto
		ctx := context.WithValue(r.Context(), "userid", userID)

		// Atualizar o contexto da requisição com o novo contexto contendo o ID do usuário
		*r = *r.WithContext(ctx)

		return true, nil
	}

	return false, errors.New("Token de autenticação inválido")
}

func extraction(r *http.Request) string {
	token := r.Header.Get("Authorization")
	log.Println(token)
	if token != "" {
		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) == 2 {
			return splitToken[1]
		}
	}
	log.Println(token)
	return ""
}

func returnKeyVerification(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado: %v", token.Header["alg"])
	}
	return []byte(secretkey), nil
}
