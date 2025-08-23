package main

import (
	"fmt"
	repl "leetcode/src/compiler-in-go/REPL"
	"os"
)

func main() {
	fmt.Println("compiler start")
	repl.Repl(os.Stdin)
}
