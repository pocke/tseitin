package eval

import (
	"fmt"
	"strings"

	"github.com/pocke/tseitin/ast"
)

type Evaluator struct {
	x string
}

type Var struct {
	Sign    bool
	Literal string
}

func (v Var) String() string {
	ret := ""
	if !v.Sign {
		ret += "!"
	}
	return ret + v.Literal
}

type Expr [][]Var

func New() *Evaluator {
	return &Evaluator{x: "A"}
}

func (ev *Evaluator) Evaluate(a ast.Expression) Expr {
	ret, _ := ev.eval(a)
	return ret
}

func (e Expr) String() string {
	ret := make([]string, 0, len(e))
	for _, v := range e {
		sslice := make([]string, 0, len(v))
		for _, x := range v {
			sslice = append(sslice, x.String())
		}
		s := strings.Join(sslice, "|")
		ret = append(ret, fmt.Sprintf("(%s)", s))
	}

	return strings.Join(ret, "&")
}

func (ev *Evaluator) eval(a ast.Expression) (Expr, string) {
	switch e := a.(type) {
	case ast.Literal:
		return Expr{}, e.Literal

	case ast.NotOpExpr:
		x := ev.Succ()
		res, y := ev.eval(e.Right)
		return append(res,
			[]Var{{Sign: false, Literal: x}, {Sign: false, Literal: y}},
			[]Var{{Sign: true, Literal: x}, {Sign: true, Literal: y}},
		), x

	case ast.BinOpExpr:
		x := ev.Succ()
		resR, y := ev.eval(e.Right)
		resL, z := ev.eval(e.Left)
		res := append(resR, resL...)

		switch e.Operator {
		case '|':
			return append(res,
				[]Var{{Sign: false, Literal: y}, {Sign: true, Literal: x}},
				[]Var{{Sign: false, Literal: z}, {Sign: true, Literal: x}},
				[]Var{{Sign: false, Literal: x}, {Sign: true, Literal: y}, {Sign: true, Literal: z}},
			), x
		case '&':
			return append(res,
				[]Var{{Sign: false, Literal: x}, {Sign: true, Literal: y}},
				[]Var{{Sign: false, Literal: x}, {Sign: true, Literal: z}},
				[]Var{{Sign: false, Literal: y}, {Sign: false, Literal: z}, {Sign: true, Literal: x}},
			), x
		default:
			panic("")
		}
	default:
		panic("")
	}
}

// A -> B -> ... Z -> AA -> AB -> ... -> AZ -> BA -> ...
func (ev *Evaluator) Succ() string {
	res := ev.x
	b := []byte(res)
	b = ev.succ(b)
	ev.x = string(b)
	return res
}

func (ev *Evaluator) succ(b []byte) []byte {
	if last := b[len(b)-1]; last != 'Z' {
		b[len(b)-1] = last + 1
		return b
	} else if len(b) == 1 {
		return []byte("AA")
	} else {
		return append(ev.succ(b[:len(b)-1]), 'A')
	}
}
