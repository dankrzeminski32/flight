package token

type TokenType string

const (
	LPAR = "("
	RPAR = ")"
	LBRACKET = "["
	RBRACKET = "]"
	LBRACE = "{"
	RBRACE = "}"
	ASSIGN = "="
	PLUS = "+"
	FUNCTION = "FUNCTION"
	ILLEGAL = "ILLEGAL"
	INTEGER = "INT"
	EOF = "EOF"
	IDENT = "IDENT"
	COMMA = ","
)

type Token struct {
	Type TokenType
	Literal string
}