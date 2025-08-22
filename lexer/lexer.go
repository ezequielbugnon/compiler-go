package lexer

const EQUAL = "="
const EOF = "EOF"

type lexer struct {
	Identifier string
	Value      string
}

type lexerImp []lexer

func (l *lexerImp) Apply(lexer string) {
	switch lexer {
	case "=":
		*l = append(*l, addLexer(EQUAL, "="))
	default:
		*l = append(*l, addLexer(EOF, "EOF"))
	}
}

func addLexer(identifier, value string) lexer {
	return lexer{
		Identifier: identifier,
		Value:      value,
	}
}
