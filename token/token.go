package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	LPAR string = "("
	RPAR string = ")"
	LBRACKET string = "["
	RBRACKET string = "]"
	ASSIGN string = "="
	PLUS string = "+"
	FUNCTION string = "FUNCTION"
	ILLEGAL string = "ILLEGAL"
	INTEGER string = "INT"
	EOF string = "EOF"
	IDENT string = "IDENT"
)