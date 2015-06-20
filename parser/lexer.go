package parser

import (
	"io"
	"text/scanner"

	"github.com/pocke/tseitin/ast"
)

type Lexer struct {
	scanner.Scanner
	result ast.Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Ident {
		if !isLower(l.TokenText()) {
			l.Error("Syntax Error")
		}
		token = LITERAL
	}
	lval.token = Token{Token: token, Literal: l.TokenText()}
	return token
}

func isLower(t string) bool {
	for _, r := range t {
		if r < 97 || 122 < r {
			return false
		}
	}
	return true
}

func (l *Lexer) Error(e string) {
	panic(e)
}

func Parse(r io.Reader) ast.Expression {
	l := new(Lexer)
	l.Init(r)
	yyParse(l)
	return l.result
}
