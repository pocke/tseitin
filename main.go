package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pocke/tseitin/parser"
)

func main() {
	r := strings.NewReader(os.Args[1])
	res := parser.Parse(r)
	fmt.Println(res)
}
