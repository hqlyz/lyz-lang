package repl

import (
	"bufio"
	"io"
	"lyz-lang/lexer"
	"lyz-lang/parser"
)

// PROMPT is a console prompt symbol
const PROMPT = ">> "

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

// Start function
func Start(in io.Reader, out io.Writer) {
	scan := bufio.NewScanner(in)
	w := bufio.NewWriter(out)

	for {
		w.WriteString(PROMPT)
		w.Flush()
		scanned := scan.Scan()
		if !scanned {
			return
		}
		line := scan.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errs()) != 0 {
			printParserErrors(w, p.Errs())
			w.Flush()
			continue
		}
		w.WriteString(program.String())
		w.WriteString("\n")
		w.Flush()
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
