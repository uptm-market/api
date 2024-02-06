package tel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// ErrInvalid erro de telefone inválido.
var ErrInvalid = errors.New("telefone inválido")

// ErrInvalidDDD erro de DDD inválido.
var ErrInvalidDDD = errors.New("DDD inválido")

// ErrInvalidNumber erro de número inválido.
var ErrInvalidNumber = errors.New("número inválido")

// Telephone representa um telefone.
type Telephone struct {
	ddd    uint8
	number uint32
}

// New cria uma instância de telefone válida.
func New(num string) (*Telephone, error) {
	ddd, number, err := captureAndValidate(num)
	if err != nil {
		return nil, err
	}
	return &Telephone{ddd, number}, nil
}

func captureAndValidate(num string) (ddd uint8, number uint32, err error) {
	if num == "" {
		return
	}
	re := regexp.MustCompile(`^\+?(55)? ?(\((\d{2})\)|(\d{2})) ?((\d)[ \.-]?)?(\d{4})[- ]?(\d{4})$`)
	if g := re.FindAllStringSubmatch(num, -1); g != nil {
		dddStr := g[0][3]
		if dddStr == "" {
			dddStr = g[0][4]
		}
		dddParsed, _ := strconv.ParseUint(dddStr, 10, 8)
		if !validateDDD(uint8(dddParsed)) {
			err = ErrInvalid
			return
		}
		ddd = uint8(dddParsed)
		numberParsed, _ := strconv.ParseUint(g[0][6]+g[0][7]+g[0][8], 10, 32)
		number = uint32(numberParsed)
		return
	}
	err = ErrInvalid
	return
}

// Empty diz se o telefone está vazio ou não
func (t *Telephone) Empty() bool {
	return t.ddd == 0 || t.number == 0
}

// DDD retorna o DDD do telefone.
func (t *Telephone) DDD() uint8 {
	return t.ddd
}

// SetDDD redefine o DDD.
func (t *Telephone) SetDDD(ddd uint8) error {
	if validateDDD(ddd) {
		t.ddd = ddd
		return nil
	}
	return ErrInvalidDDD
}

func validateDDD(ddd uint8) bool {
	validDDDs := []uint8{11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 24, 27, 28, 31, 32, 33, 34, 35, 37, 38, 41, 42, 43, 44, 45, 46, 47, 48, 49, 51, 53, 54, 55, 61, 62, 63, 64, 65, 66, 67, 68, 69, 71, 73, 74, 75, 77, 79, 81, 82, 83, 84, 85, 86, 87, 88, 89, 91, 92, 93, 94, 95, 96, 97, 98, 99}
	i := sort.Search(len(validDDDs), func(i int) bool {
		return validDDDs[i] >= ddd
	})
	return i > -1
}

// Number retorna o número do telefone.
func (t *Telephone) Number() uint32 {
	return t.number
}

// SetNumber redefine o número do telefone.
func (t *Telephone) SetNumber(num uint32) error {
	if num <= 999999999 {
		t.number = num
		return nil
	}
	return ErrInvalidNumber
}

// FullNumber retorna o número completo sem formatação.
func (t *Telephone) FullNumber() string {
	if t.Empty() {
		return ""
	}
	return fmt.Sprintf("%d%d", t.ddd, t.number)
}

// String retorna o número completo e formatado.
func (t Telephone) String() string {
	if t.Empty() {
		return ""
	}
	s := t.FullNumber()
	var r strings.Builder
	r.WriteString("(" + s[:2] + ") ")
	if len(s) > 10 {
		r.WriteString(s[2:3] + " ")
		r.WriteString(s[3:7] + "-" + s[7:])
	} else {
		r.WriteString(s[2:6] + "-" + s[6:])
	}
	return r.String()
}

// Value implementa a interface Valuer de database/sql.
func (t Telephone) Value() (driver.Value, error) {
	return t.FullNumber(), nil
}

// Scan implementa a interface Scanner para conversão no banco de dados.
func (t *Telephone) Scan(src interface{}) error {
	if bv, err := driver.String.ConvertValue(src); err == nil {
		if v, ok := bv.(string); ok {
			tel, err := New(v)
			if err != nil {
				return err
			}
			*t = *tel
			return nil
		}

		if v, ok := bv.([]byte); ok {
			tel, err := New(string(v))
			if err != nil {
				return err
			}
			*t = *tel
			return nil
		}
	}
	return errors.New("failed to scan Telephone")
}

// MarshalJSON implementa a interface json.Marshaler.
func (t Telephone) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.FullNumber())
}

// UnmarshalJSON implementa a interface json.Unmarshaler.
func (t *Telephone) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	tel, err := New(v)
	if err != nil {
		return err
	}
	t.SetDDD(tel.DDD())
	t.SetNumber(tel.Number())
	return nil
}
