package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Numeez/interpreter/lexer"
	"github.com/Numeez/interpreter/token"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out,prompt)
		if !scanner.Scan(){
			return
		}
		line:= scanner.Text()
		l:=lexer.New(line)
		for tok:=l.NextToken();tok.Type!=token.EOF;tok=l.NextToken(){
			fmt.Fprintf(out, "%+v\n", tok)
		}

	}
}
