package eval

import (
	"fmt"
	"strings"
)

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
