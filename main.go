package main

import (
	"fmt"
	"os"

	"github.com/cmcki85/yeet-interpreter/repl"
)

func main() {
    fmt.Println("This is YEET, a new language designed for gen-z devs!")
    fmt.Println("Type in some commands to get started\n")
    repl.Start(os.Stdin, os.Stdout)
}

