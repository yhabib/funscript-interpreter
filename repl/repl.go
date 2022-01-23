package repl

import (
	"bufio"
	"fmt"
	"funscript-interpreter/lexer"
	"funscript-interpreter/token"
	"io"
)

const PROMPT = "❤️"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("%s ", PROMPT)
		input := scanner.Scan()
		if !input {
			return
		}
		line := scanner.Text()
		lexed := lexer.New(line)
		for tok := lexed.NextToken(); tok.Type != token.EOF; tok = lexed.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
