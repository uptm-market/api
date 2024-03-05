package entity

import (
	"time"

	"go.mod/doc"
	"go.mod/tel"
)

type User struct {
	ID                    uint          `json:"id,omitempty"`
	Email                 string        `json:"email"`
	EmailVerificationKey  string        `json:"-"`
	EmailVerificationTime *time.Time    `json:"-"`
	Password              string        `json:"password"`
	Name                  string        `json:"name"`
	CellPhone             tel.Telephone `json:"cellPhone"`
	City                  string        `json:"city,omitempty"`
	State                 string        `json:"state,omitempty"`
	ZipCode               string        `json:"zipCode,omitempty"`
	District              string        `json:"district,omitempty"`
	Street                string        `json:"street,omitempty"`
	StreetNumber          string        `json:"streetNumber,omitempty"`
	Complement            string        `json:"complement,omitempty"`
	Birthdate             string        `json:"birthdate,omitempty"`
	Gender                string        `json:"gender,omitempty"`
	CPF                   *doc.CPF      `json:"cpf"`
	Source                uint8         `json:"source,omitempty"`
	RecordDate            time.Time     `json:"recordDate,omitempty"`
	Level                 uint8         `json:"-"`
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
