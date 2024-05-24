package core

import (
	"context"
	"log"

	"go.mod/db"
	"go.mod/entity"
	"go.mod/middleware"

	"go.mod/rest"
)

type UserManager struct {
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (um *UserManager) User(ctx context.Context, id string) (*entity.UserInfoView, error) {
	userinfo, err := db.ReturnUserById(ctx, id)
	if err != nil {
		return nil, rest.LogError(err)
	}
	return userinfo, nil
}
func (um *UserManager) CreateUser(ctx context.Context, user entity.UserCreations) error {

	verify, err := db.VerifyUserExists(ctx, user.Email)
	if err != nil {
		return rest.LogError(err)
	}
	if verify {
		return &rest.Error{Status: 400, Message: "Esse e-mail já esta em uso na plataforma", Code: "aleready_exists"}
	}

	err = db.Create(ctx, user)
	if err != nil {
		return rest.LogError(err, "um.CreateUser db.Create")
	}
	return nil
}

func (*UserManager) Login(ctx context.Context, email, password string) (string, error) {
	// Implemente a lógica para verificar as credenciais do usuário no banco de dados.
	// Aqui, estou usando uma função fictícia chamada VerifyCredentials como exemplo.
	user, err := db.VerifyCredentials(ctx, email, password)
	if err != nil {
		return "", &rest.Error{Status: 400, Code: "invalid_user", Message: "Usuario não existe na plataforma"}
	}
	if user == nil || user.ID == 0 {
		return "", &rest.Error{Status: 400, Code: "invalid_user", Message: "Usuario não existe na plataforma"}
	}
	body := &entity.User{
		ID:    user.ID,
		Email: user.Email,
	}
	token, err := middleware.GenerateToken(body)
	if err != nil {
		return "", rest.LogError(err, "middleware.generatetoken")
	}
	return token, nil
}

func (m *UserManager) GetMeInfo(ctx context.Context) (*entity.ReturnUserInfo, error) {
	info, err := db.ReturnInfoMe(ctx, ctx.Value("userID").(string))
	if err != nil {
		log.Println("user.GetMeInfo db.ReturnInfoMe error:", err)
		return nil, err
	}
	return info, nil

}
