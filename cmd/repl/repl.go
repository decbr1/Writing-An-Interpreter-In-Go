// ./cmd/repl/repl.go

package repl

import (
	"WritingAnInterpreterInGo/cmd/lexer"
	"WritingAnInterpreterInGo/cmd/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

// Start begins the REPL (Read Eval Print Loop). The REPL reads input, sends it to the
// interpreter for evaluation, prints the result/output of the interpreter and starts again.
// Read, Eval, Print, Loop. Currently we only tokenize Monkey source code and print the tokens.
// Later on, we will expand this and add parsing and evaluation to it.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		// Print the prompt
		fmt.Fprintf(out, PROMPT)
		// Read from the input source until encountering a newline
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		// Take the just-read line and pass it to an instance of our lexer
		line := scanner.Text()
		l := lexer.New(line)

		// Print all the tokens the lexer gives us until we encounter EOF
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
