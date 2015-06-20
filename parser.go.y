%{
package main

import (
  "fmt"
  "text/scanner"
  "os"
  "strings"
)

type Expression interface{}

type Token struct {
  token int
  literal string
}

type Literal struct {
  literal string
}

type NotOpExpr struct {
	operator rune
	right Expression
}

type BinOpExpr struct {
  left Expression
  operator rune
  right Expression
}

type ParenExpr struct {
	subExpr Expression
}

%}


%union{
  token Token
  expr Expression
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
		$$ = Literal{literal: $1.literal}
	}
	| and_expr
	| or_expr
	| not_expr
	| paren_expr

and_expr
	: expr '&' expr
	{
		$$ = BinOpExpr{left: $1, operator: '&', right: $3}
	}

or_expr
	: expr '|' expr
	{
		$$ = BinOpExpr{left: $1, operator: '|', right: $3}
	}

not_expr
	: '!' expr
	{
		$$ = NotOpExpr{operator: '!', right: $2}
	}

paren_expr
	: '(' expr ')'
	{
		$$ = ParenExpr{subExpr: $2}
	}

%%

type Lexer struct {
	scanner.Scanner
	result Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Ident {
		token = LITERAL
	}
	lval.token = Token{token: token, literal: l.TokenText()}
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
