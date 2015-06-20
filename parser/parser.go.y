%{
package main

import (
  "fmt"
  "text/scanner"
  "os"
  "strings"
  "github.com/pocke/tseitin/ast"
)

type Token struct {
	Token   int
	Literal string
}

%}


%union{
  token Token
  expr ast.Expression
}

%type<expr> program
%type<expr> expr and_expr or_expr not_expr paren_expr
%token<token> LITERAL

%left '&' '|'
%right '!'

%%

program
  : expr
  {
    $$ = $1
    yylex.(*Lexer).result = $$
  }

expr
	: LITERAL
	{
		$$ = ast.Literal{Literal: $1.Literal}
	}
	| and_expr
	| or_expr
	| not_expr
	| paren_expr

and_expr
	: expr '&' expr
	{
		$$ = ast.BinOpExpr{Left: $1, Operator: '&', Right: $3}
	}

or_expr
	: expr '|' expr
	{
		$$ = ast.BinOpExpr{Left: $1, Operator: '|', Right: $3}
	}

not_expr
	: '!' expr
	{
		$$ = ast.NotOpExpr{Operator: '!', Right: $2}
	}

paren_expr
	: '(' expr ')'
	{
		$$ = ast.ParenExpr{SubExpr: $2}
	}

%%

type Lexer struct {
	scanner.Scanner
	result ast.Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Ident {
		token = LITERAL
	}
	lval.token = Token{Token: token, Literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}

func main() {
	l := new(Lexer)
	l.Init(strings.NewReader(os.Args[1]))
	yyParse(l)
	fmt.Printf("%#v\n", l.result)
}
