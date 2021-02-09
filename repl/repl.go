package repl

import (
	"bufio"
	"fmt"
	"io"
	"lyz-lang/lexer"
	"lyz-lang/token"
)

// PROMPT is a console prompt symbol
const PROMPT = ">> "

// Start function
func Start(in io.Reader, out io.Writer) {
	scan := bufio.NewScanner(in)
	w := bufio.NewWriter(out)

	for {
		w.WriteString(PROMPT)
		w.Flush()
		if !scan.Scan() {
			return
		}
		l := lexer.New(scan.Text())

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			w.WriteString(fmt.Sprintf("%+v\n", t))
		}
		w.Flush()
	}
}
