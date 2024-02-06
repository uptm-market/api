package doc

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const cpfLen = 11

// CPF representa um número de CPF.
type CPF struct {
	number [cpfLen]uint8
}

// NewCPF cria um novo CPF sem validar `number`.
func NewCPF(number string) *CPF {
	number = strings.Map(func(r rune) rune {
		if r < '0' || r > '9' {
			return -1
		}
		return r
	}, number)
	if number == "" {
		return &CPF{}
	}
	if len(number) > cpfLen {
		number = number[:cpfLen]
	}
	return &CPF{prepareCPFNumber(number)}
}

// ParseCPF cria um CPF a partir de uma string.
// Não valida os dígitos verificadores.
func ParseCPF(number string) (*CPF, error) {
	if number == "" {
		return &CPF{}, nil
	}
	pattern, err := regexp.Compile(cpfFormat)
	if err != nil {
		return nil, fmt.Errorf("cpf: %w", err)
	}
	if !pattern.MatchString(number) {
		return nil, errors.New("cpf: invalid format")
	}
	groups := pattern.FindStringSubmatch(number)
	rawNumber := groups[1] + groups[2] + groups[3] + groups[4]
	return &CPF{number: prepareCPFNumber(rawNumber)}, nil
}

func prepareCPFNumber(s string) (result [cpfLen]uint8) {
	for i, r := range s {
		result[i] = uint8(r - '0')
	}
	return result
}

// IsValid verifica se o(s) dígito(s) verificador(es) é(são) válido(s).
func (c CPF) IsValid() bool {
	sum := 0
	for i := 0; i < len(c.number)-2; i++ {
		sum += int(c.number[i]) * (10 - i)
	}
	d1 := uint8(sum % 11)
	if d1 < 2 {
		d1 = 0
	} else {
		d1 = 11 - d1
	}

	sum = 0
	for i := 0; i < len(c.number)-2; i++ {
		sum += int(c.number[i]) * (11 - i)
	}
	sum += int(d1) * 2
	d2 := uint8(sum % 11)
	if d2 < 2 {
		d2 = 0
	} else {
		d2 = 11 - d2
	}
	return d1 == c.number[9] && d2 == c.number[10]
}

// MarshalJSON codifica o CPF para JSON.
func (c CPF) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.FullNumber(false))
}

// UnmarshalJSON decodifica o CPF do JSON.
func (c *CPF) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if cpf, err := ParseCPF(s); err == nil {
		c.number = cpf.number
	} else {
		return err
	}
	return nil
}

// Value implementa a interface Valuer de database/sql.
func (c CPF) Value() (driver.Value, error) {
	return c.FullNumber(false), nil
}

// Scan implementa a interface Scanner para conversão no banco de dados.
func (c *CPF) Scan(src interface{}) error {
	var cpf *CPF
	switch val := src.(type) {
	case []byte:
		cpf = NewCPF(string(val))
	case string:
		cpf = NewCPF(val)
	case nil:
		c.number = [cpfLen]uint8{}
		return nil
	default:
		return errors.New("cpf: failed to scan")
	}
	c.number = cpf.number
	return nil
}

// IsEmpty retorna true se o CPF estiver em branco, caso contrário, false.
func (c CPF) IsEmpty() bool {
	return c.number == [cpfLen]uint8{}
}

// FullNumber retorna o número do CPF com dígito verificador e sem formatação.
func (c CPF) FullNumber(formated bool) string {
	if c.IsEmpty() {
		return ""
	}
	var result strings.Builder
	for i, n := range c.number {
		result.WriteByte(byte('0' + n))
		if formated {
			switch i {
			case 2, 5:
				result.WriteRune('.')
			case 8:
				result.WriteRune('-')
			}
		}
	}
	return result.String()
}

func (c CPF) String() string {
	return c.FullNumber(true)
}

// Type retorna o tipo do documento.
// Sempre "cpf".
func (CPF) Type() string {
	return "cpf"
}
