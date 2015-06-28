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

	fmt.Println(tseitin.String())
}
