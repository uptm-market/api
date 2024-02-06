package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mod/entity"
)

// SecretKey é a chave secreta usada para assinar o token.
var SecretKey = []byte("SeCuCGdto5Bu7notHb5EvEV0I27pW78PePRLznFG6iV5iHAW2RA9BYDTSj7sotvMuWR63TT1LrvD+yKzyyDKFg==")

// GenerateToken gera um token JWT para o usuário com informações adicionais.
func GenerateToken(user *entity.User) (string, error) {
	// Criação do payload do token com as informações desejadas.
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"level": user.Level,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expira em 24 horas (ajuste conforme necessário).
	}

	// Criação do token com os claims e a assinatura usando a chave secreta.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
