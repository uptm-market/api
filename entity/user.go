package entity

import (
	"time"

	"go.mod/doc"
	"go.mod/tel"
)

type User struct {
	ID                    uint          `json:"id,omitempty"`
	Email                 string        `json:"email"`
	EmailVerificationKey  string        `json:"emailVerificationKey"`
	EmailVerificationTime *time.Time    `json:"emailVerificationTime"`
	Password              string        `json:"password"`
	Name                  string        `json:"name"`
	CellPhone             tel.Telephone `json:"cellPhone"`
	City                  string        `json:"city"`
	State                 string        `json:"state"`
	ZipCode               string        `json:"zipCode"`
	District              string        `json:"district"`
	Street                string        `json:"street"`
	StreetNumber          string        `json:"streetNumber"`
	Complement            string        `json:"complement"`
	Birthdate             string        `json:"birthdate"`
	Gender                string        `json:"gender"`
	CPF                   *doc.CPF      `json:"cpf"`
	Source                uint8         `json:"source"`
	RecordDate            time.Time     `json:"recordDate"`
	Level                 uint8         `json:"-"`
}

type UserCreations struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserInfoView struct {
	ID           uint          `json:"id,omitempty"`
	Email        string        `json:"email"`
	Name         string        `json:"name"`
	CellPhone    tel.Telephone `json:"cellPhone"`
	City         string        `json:"city,omitempty"`
	State        string        `json:"state,omitempty"`
	ZipCode      string        `json:"zipCode,omitempty"`
	District     string        `json:"district,omitempty"`
	Street       string        `json:"street,omitempty"`
	StreetNumber string        `json:"streetNumber,omitempty"`
	Complement   string        `json:"complement,omitempty"`
	Birthdate    string        `json:"birthdate,omitempty"`
	Gender       string        `json:"gender,omitempty"`
	CPF          *doc.CPF      `json:"cpf"`
	Source       uint8         `json:"source,omitempty"`
	RecordDate   time.Time     `json:"recordDate,omitempty"`
	Level        uint8         `json:"-"`
}
type ReturnUserInfo struct {
	ID        uint          `json:"id,omitempty"`
	Email     string        `json:"email"`
	Name      string        `json:"name"`
	CellPhone tel.Telephone `json:"cellPhone"`
	City      string        `json:"city,omitempty"`
	State     string        `json:"state,omitempty"`
	Level     uint8         `json:"-"`
}

type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
