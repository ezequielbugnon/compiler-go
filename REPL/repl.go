package repl

import (
	"bufio"
	"fmt"
	"io"
	"leetcode/src/compiler-in-go/lexer"
	"leetcode/src/compiler-in-go/token"
)

const PROMPT = ">> "

func Repl(in io.Reader) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}

}
