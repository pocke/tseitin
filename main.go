package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pocke/tseitin/eval"
	"github.com/pocke/tseitin/parser"
)

func main() {
	r := strings.NewReader(os.Args[1])
	res := parser.Parse(r)
	ev := eval.New()
	tseitin := ev.Evaluate(res)

	s := ""
	for _, v := range tseitin {
		s += "(" + v + ")&"
	}

	b := []byte(s)
	s = string(b[:len(b)-1])
	fmt.Println(s)
}
