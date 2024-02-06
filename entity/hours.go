package entity

type Hours struct {
	Hour           string `json:"hour,omitempty"`
	Professional   string `json:"profissional,omitempty"`
	Local          string `json:"local,omitempty"`
	UserId         string `json:"userId,omitempty"`
	UserName       string `json:"userName,omitempty"`
	ProfessionalId string `json:"profissionalId,omitempty"`
	CreatedTime    string `json:"created_time,omitempty"`
}
