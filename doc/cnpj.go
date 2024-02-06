package doc

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const cnpjLen = 14

// CNPJ representa um número de CNPJ.
type CNPJ struct {
	number [cnpjLen]uint8
}

// NewCNPJ cria um novo CNPJ sem validar `number`.
func NewCNPJ(number string) *CNPJ {
	number = strings.Map(func(r rune) rune {
		if r < '0' || r > '9' {
			return -1
		}
		return r
	}, number)
	if number == "" {
		return &CNPJ{}
	}
	if len(number) > cnpjLen {
		number = number[:cnpjLen]
	}
	return &CNPJ{prepareCNPJNumber(number)}
}

// ParseCNPJ cria um CNPJ a partir de uma string validando o formato.
// Mas não valida os dígitos verificadores.
func ParseCNPJ(number string) (*CNPJ, error) {
	if number == "" {
		return &CNPJ{}, nil
	}
	pattern, err := regexp.Compile(cnpjFormat)
	if err != nil {
		return nil, fmt.Errorf("cnpj: %w", err)
	}
	if !pattern.MatchString(number) {
		return nil, errors.New("cnpj: invalid format")
	}
	groups := pattern.FindStringSubmatch(number)
	rawNumber := groups[1] + groups[2] + groups[3] + groups[4] + groups[5]
	return &CNPJ{number: prepareCNPJNumber(rawNumber)}, nil
}

// IsValid verifica se o(s) dígito(s) verificador(es) é(são) válido(s).
func (c CNPJ) IsValid() bool {
	sum := 0
	l := len(c.number)
	for i := 0; i < l-2; i++ {
		sum += int(c.number[i]) * ((l-3-i)%8 + 2)
	}
	d1 := uint8(sum % 11)
	if d1 < 2 {
		d1 = 0
	} else {
		d1 = 11 - d1
	}

	sum = 0
	for i := 0; i < l-2; i++ {
		sum += int(c.number[i]) * ((l-2-i)%8 + 2)
	}
	sum += int(d1) * 2
	d2 := uint8(sum % 11)
	if d2 < 2 {
		d2 = 0
	} else {
		d2 = 11 - d2
	}
	return d1 == c.number[12] && d2 == c.number[13]
}

func prepareCNPJNumber(s string) (result [cnpjLen]uint8) {
	for i, r := range s {
		result[i] = uint8(r - '0')
	}
	return result
}

// MarshalJSON codifica o CNPJ para JSON.
func (c CNPJ) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.FullNumber(false))
}

// UnmarshalJSON decodifica o CNPJ do JSON.
func (c *CNPJ) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if cnpj, err := ParseCNPJ(s); err == nil {
		c.number = cnpj.number
	} else {
		return err
	}
	return nil
}

// Value implementa a interface Valuer de database/sql.
func (c CNPJ) Value() (driver.Value, error) {
	return c.FullNumber(false), nil
}

// Scan implementa a interface Scanner para conversão no banco de dados.
func (c *CNPJ) Scan(src interface{}) error {
	var cnpj *CNPJ
	switch val := src.(type) {
	case []byte:
		cnpj = NewCNPJ(string(val))
	case string:
		cnpj = NewCNPJ(val)
	case nil:
		c.number = [cnpjLen]uint8{}
		return nil
	default:
		return errors.New("cnpj: failed to scan")
	}
	c.number = cnpj.number
	return nil
}

// IsEmpty retorna true se o CNPJ estiver em branco, caso contrário, false.
func (c CNPJ) IsEmpty() bool {
	return c.number == [cnpjLen]uint8{}
}

// FullNumber retorna o número do CNPJ com dígito verificador e sem formatação.
func (c CNPJ) FullNumber(formated bool) string {
	if c.IsEmpty() {
		return ""
	}
	var result strings.Builder
	for i, n := range c.number {
		result.WriteByte(byte('0' + n))
		if formated {
			switch i {
			case 1, 4:
				result.WriteRune('.')
			case 7:
				result.WriteRune('/')
			case 11:
				result.WriteRune('-')
			}
		}
	}
	return result.String()
}

// String retorna o número do CNPJ formatado.
func (c CNPJ) String() string {
	return c.FullNumber(true)
}

// Type retorna o tipo do documento.
// Sempre "cnpj".
func (CNPJ) Type() string {
	return "cnpj"
}
