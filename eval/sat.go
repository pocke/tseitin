package eval

import (
	"regexp"

	"github.com/pocke/go-minisat"
)

func (e Expr) Solve() (bool, map[string]bool) {
	s := minisat.NewSolver()

	table := make(map[string]*minisat.Var)

	for _, clause := range e {
		satc := make([]*minisat.Var, 0, len(clause))
		for _, v := range clause {
			satv, exist := table[v.Literal]
			if !exist {
				satv = s.NewVar()
				table[v.Literal] = satv
			}

			if !v.Sign {
				satv = satv.Not()
			}
			satc = append(satc, satv)
		}
		s.AddClause(satc...)
	}

	isSat := s.Solve()

	if !isSat {
		return isSat, nil
	}

	retTable := make(map[string]bool)
	re := regexp.MustCompile(`^[a-z]+$`)
	for lit, satv := range table {
		if !re.MatchString(lit) {
			continue
		}
		b, _ := s.ModelValue(satv)
		retTable[lit] = b
	}

	return isSat, retTable
}
