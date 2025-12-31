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

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Monkey programming language.\n", user.Username)
	fmt.Printf("This is the REPL. Feel free to type some commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}
