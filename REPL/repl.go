package repl

import (
	"bufio"
	"fmt"
	"io"
)

func Repl(in io.Reader) {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		fmt.Println("working", scanner.Text())
	}
}
