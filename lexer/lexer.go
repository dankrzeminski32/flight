package lexer

type Lexer struct {
	input string 
	position int 
	nextposition int
	cc byte // current character
}

func New(input string) (*Lexer) {
	return &Lexer{input: input}
}