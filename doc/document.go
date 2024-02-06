package doc

import "regexp"

// Document representa um documento.
type Document interface {
	IsValid() bool
	FullNumber(formated bool) string
	Type() string
}

const (
	cpfFormat  = `^(\d{3})\.?(\d{3})\.?(\d{3})-?(\d{2})$`
	cnpjFormat = `^(\d{2})\.?(\d{3})\.?(\d{3})/?(\d{4})-?(\d{2})$`
)

// Type retorna o tipo de documento de acordo com o número.
// Retorna "cpf", "cnpj" ou "".
func Type(str string) string {
	if matched, _ := regexp.MatchString(cpfFormat, str); matched {
		return "cpf"
	}
	if matched, _ := regexp.MatchString(cnpjFormat, str); matched {
		return "cnpj"
	}
	return ""
}

// IsDocumentNumber diz se é um documento válido.
func IsDocumentNumber(str string) bool {
	return Type(str) != ""
}
