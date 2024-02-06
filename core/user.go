package core

import (
	"context"
	"fmt"

	"go.mod/db"
	"go.mod/entity"
	"go.mod/middleware"
)

type UserManager struct {
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (*UserManager) User(ctx context.Context) (*entity.User, error) {
	fmt.Println("teste mateus henrique")
	return nil, nil
}
func (*UserManager) CreateUser(ctx context.Context, user entity.User) error {

	return nil
}

func (*UserManager) Login(ctx context.Context, email, password string) (string, error) {
	// Implemente a lógica para verificar as credenciais do usuário no banco de dados.
	// Aqui, estou usando uma função fictícia chamada VerifyCredentials como exemplo.
	user, err := db.VerifyCredentials(ctx, email, password)
	if err != nil {
		return "", err
	}
	body := &entity.User{
		ID:    user.ID,
		Email: user.Email,
	}
	token, err := middleware.GenerateToken(body)
	if err != nil {
		return "", err
	}
	return token, nil
}
