package db

import (
	"context"
	"time"

	infradb "go.mod/connect"
	"go.mod/entity"
	"go.mod/rest"
)

func Create(ctx context.Context, user entity.User) error {
	query := `
        INSERT INTO users (
            email,
            email_verification_key,
            email_verification_time,
            password,
            name,
            birth_date,
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
            level
        ) VALUES (
            $1, 'verification', NOW(), $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
        )`

	_, err := infradb.Get().QueryContext(ctx, query,
		user.Email,
		user.Password,
		user.Name,
		user.Birthdate,
		user.Gender,
		user.CPF.String(), // Assuming CPF is a custom type with a String method
		user.CellPhone,
		user.ZipCode,
		user.State,
		user.City,
		user.District,
		user.Street,
		user.StreetNumber,
		user.Complement,
		user.Level,
	)

	return err
}

func VerifyCredentials(ctx context.Context, email, password string) (*entity.User, error) {
	user := &entity.User{
		Email:    email,
		Password: password,
	}

	query := "SELECT id, email, level FROM users WHERE email = $1 AND password = $2"
	rows, err := infradb.Get().QueryContext(ctx, query, email, password)
	if err != nil {
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
	birth_date,
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
