package token

type TokenType string
type Token struct {
	Type  TokenType
	Value string
}

var keyWords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupreservedWords(identifier string) TokenType {
	if token, ok := keyWords[identifier]; ok {
		return token
	}
	return IDENT
}
