%{
package main

import (
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
