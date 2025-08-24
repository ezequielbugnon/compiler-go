package lexer

const EQUAL = "="
const PLUS = "+"
const MINUS = "-"
const LET = "LET"
const EOF = "EOF"
const IF = "IF"
const ELSE = "ELSE"
const LEFTBRACKET = "("
const RIGHTBRACKET = ")"

type lexer struct {
	Identifier string
	Value      string
}

type lexerImp []lexer

func New() *lexerImp {
	return &lexerImp{}
}

func (l *lexerImp) Apply(lexer string) {

	switch lexer {
	case "=":
		*l = append(*l, addLexer(EQUAL, "="))
	case "+":
		*l = append(*l, addLexer(PLUS, "+"))
	case "-":
		*l = append(*l, addLexer(MINUS, "-"))
	case "(":
		*l = append(*l, addLexer(LEFTBRACKET, "("))
	case "":
		*l = append(*l, addLexer(RIGHTBRACKET, ")"))
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
