package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cmcki85/yeet-interpreter/lexer"
	"github.com/cmcki85/yeet-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Print(PROMPT)
        scanned := scanner.Scan()
        if (!scanned) {
            return
        }

        line := scanner.Text()
        lexer := lexer.New(line)

        for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
    }
}
