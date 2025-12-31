// Writing an Interpreter in Go
// a book by Thorsten Ball
// https://interpreterbook.com/

// ./cmd/main.go

package main

import (
	"WritingAnInterpreterInGo/cmd/repl"
	"fmt"
	"os"
	"os/user"
)

// main welcomes the user to the Monkey REPL and starts it.
// The REPL (Read Eval Print Loop) reads input, sends it to the interpreter
// for evaluation, prints the result/output of the interpreter and starts again.
func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Monkey programming language.\n", user.Username)
	fmt.Printf("This is the REPL. Feel free to type some commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}
