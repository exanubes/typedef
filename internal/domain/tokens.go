package domain

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL  = "ILLEGAL"
	EOF      = "EOF"
	IDENT    = "IDENT"
	NUMBER   = "NUMBER"
	STRING   = "STRING"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	NULL     = "NULL"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
	COMMA    = "," //DELIMITER?
	COLON    = ":" //ASSIGN?
)

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

func NewToken(token_type TokenType, character byte) Token {
	return Token{Type: token_type, Literal: string(character)}
}

func LookupIdentifier(val string) TokenType {
	if value, exists := keywords[val]; exists {
		return value
	}
	return IDENT
}
