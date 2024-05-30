package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	infradb "go.mod/connect"
	"go.mod/entity"
	"go.mod/rest"
)

func Create(ctx context.Context, user entity.UserCreations) error {
	query := `
	INSERT INTO users (
		email,
		password,
		name,
		level
	) VALUES (
		$1,  $2, $3, $4
	)
	`

	_, err := infradb.Get().QueryContext(ctx, query,
		user.Email,
		user.Password,
		user.Name,
		1,
	)

	return err
}

func Update(ctx context.Context, user entity.UserUpdated, id uint) error {
	query := `
	UPDATE users
	SET
	email = $2,
		cpf = $3
		
	WHERE
	id = $1
		
	`

	_, err := infradb.Get().ExecContext(ctx, query,
		id,
		user.Email,
		user.CPF,
	)

	return err
}

func VerifyCredentials(ctx context.Context, email, password string) (*entity.User, error) {
	user := &entity.User{
		Email:    email,
		Password: password,
	}
	fmt.Println("testando now")
	query := "SELECT id, email, level FROM users WHERE email = $1 AND password = $2"
	rows, err := infradb.Get().QueryContext(ctx, query, email, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &rest.Error{Status: 400, Code: "user_not_found", Message: "User not found. Login declined."}
		}
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Level /* ... outros campos ... */)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, &rest.Error{Status: 400, Code: "user_not_found", Message: "User not found login declined."}
	}

	return user, nil
}

func ReturnUserById(ctx context.Context, id string) (*entity.UserInfoView, error) {
	var array []entity.UserInfoView
	rows, err := infradb.Get().QueryContext(ctx, `
	select 
	email,
	name,
	birthdate,
	gender,
	cpf,
	cell_phone,
	zip_code,
	state,
	city,
	district,
	street,
	street_number,
	complement,
	level from users where id =$1;
	`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user entity.UserInfoView
		err := rows.Scan(
			&user.Email,
			&user.Name,
			&user.Birthdate,
			&user.Gender,
			&user.CPF,
			&user.CellPhone,
			&user.ZipCode,
			&user.State,
			&user.City,
			&user.District,
			&user.Street,
			&user.StreetNumber,
			&user.Complement,
			&user.Level,
		)
		if err != nil {
			return nil, err
		}
		array = append(array, user)
	}

	return &array[0], nil
}

type User struct {
	ID                    int
	EmailVerificationTime time.Time
	CreationTime          time.Time
}

func VerificationTimeUser(ctx context.Context, userID string) (bool, error) {
	var user User
	err := infradb.Get().QueryRowContext(ctx, "SELECT id, email_verification_time, creation_time FROM users WHERE id = $1", userID).Scan(&user.ID, &user.EmailVerificationTime, &user.CreationTime)
	if err != nil {
		return false, err
	}

	// Verificar se passou 1 mês desde a criação do usuário
	if time.Since(user.CreationTime) >= (30 * 24 * time.Hour) {
		return true, nil
	}

	return false, nil
}

func ReturnInfoMe(ctx context.Context, Id uint) (*entity.ReturnUserInfo, error) {
	var data entity.ReturnUserInfo
	err := infradb.Get().QueryRowContext(ctx, `select id, name, email,cpf ,level from users where id =$1`, Id).Scan(&data.ID, &data.Name, &data.Email, &data.CPF, &data.Level)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func ReturnPassword(ctx context.Context, Id uint) (*string, error) {
	var pass string
	err := infradb.Get().QueryRowContext(ctx, `select password from users where id =$1`, Id).Scan(&pass)
	if err != nil {
		return nil, err
	}
	return &pass, nil
}

func UpdatedPassword(ctx context.Context, data entity.UpdatePassword, Id uint) error {
	log.Println("final", data)
	log.Println("id", Id)
	log.Println("data.new", data.NewPassword)

	_, err := infradb.Get().ExecContext(ctx, `
	UPDATE users
	SET
		password = $2
	WHERE
		id = $1
	`, Id, data.NewPassword)
	if err != nil {
		return err
	}
	return nil
}

func VerifyUserExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	err := infradb.Get().QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
