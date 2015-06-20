package eval

import (
	"fmt"

	"github.com/pocke/tseitin/ast"
)

type Evaluator struct {
	x string
}

func New() *Evaluator {
	return &Evaluator{x: "A"}
}

func (ev *Evaluator) Evaluate(a ast.Expression) []string {
	res, _ := ev.eval(a)
	return res
}

func (ev *Evaluator) eval(a ast.Expression) ([]string, string) {
	switch e := a.(type) {
	case ast.Literal:
		return []string{}, e.Literal

	case ast.NotOpExpr:
		x := ev.Succ()
		res, y := ev.eval(e.Right)
		return append(res,
			fmt.Sprintf("!%s|!%s", x, y),
			fmt.Sprintf("%s|%s", x, y),
		), x

	case ast.BinOpExpr:
		x := ev.Succ()
		resR, y := ev.eval(e.Right)
		resL, z := ev.eval(e.Left)
		res := append(resR, resL...)

		switch e.Operator {
		case '|':
			return append(res,
				fmt.Sprintf("!%s|%s", y, x),
				fmt.Sprintf("!%s|%s", z, x),
				fmt.Sprintf("!%s|%s|%s", x, y, z),
			), x
		case '&':
			return append(res,
				fmt.Sprintf("!%s|%s", x, y),
				fmt.Sprintf("!%s|%s", x, z),
				fmt.Sprintf("!%s|!%s|%s", y, z, x),
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
