package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// AuthMiddlewareWithClaims é um middleware que verifica a autenticidade do token e extrai as informações.
func AuthMiddlewareWithClaims(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtenha o token do cabeçalho de autorização.
		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Parse do token com a chave secreta.
		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			// Verifique se o método de assinatura é válido.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inválido")
			}
			return SecretKey, nil
		})

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Verifique se o token é válido.
		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extraia as claims do token.
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Erro ao obter informações do token", http.StatusInternalServerError)
			return
		}

		// Agora você tem acesso às informações do usuário no mapa 'claims'.
		// Exemplo: id := claims["id"].(uint)
		//          email := claims["email"].(string)
		//          level := claims["level"].(uint8)

		// Adicione as informações do usuário ao contexto para uso posterior se necessário.
		ctx := context.WithValue(r.Context(), "userID", claims["id"])
		ctx = context.WithValue(ctx, "userEmail", claims["email"])
		ctx = context.WithValue(ctx, "userLevel", claims["level"])

		// Passe para o próximo middleware ou manipulador com o contexto atualizado.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
