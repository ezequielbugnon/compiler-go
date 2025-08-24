package repl

import (
	"bufio"
	"fmt"
	"io"
	"leetcode/src/compiler-in-go/lexer"
)

func Repl(in io.Reader) {
	scanner := bufio.NewScanner(in)
	lexer := lexer.New()

	if scanner.Scan() {
		fmt.Println("working", scanner.Text())
		lexer.Apply(scanner.Text())
	}

	fmt.Println("lexer", lexer)
}
