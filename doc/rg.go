package doc

// RG representa o documento de identidade (Registro Geral).
type RG struct {
	Number string
	Issuer string
	State  string
}

// IsValid diz se o documento é válido.
func (RG) IsValid() bool {
	return true
}

// FullNumber retorna o número do documento.
func (rg RG) FullNumber(formated bool) string {
	return rg.Number
}

// Type retorna o tipo do documento.
// Sempre "rg".
func (RG) Type() string {
	return "rg"
}
