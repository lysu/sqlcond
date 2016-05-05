package sql

import (
	"strings"
	"github.com/lysu/sqlcond"
)

type SqlExpVisitor struct {
	currConds []string
	currArgs  []interface{}
	andSql    string
	andArgs   []interface{}
	orSql     string
	orArgs    []interface{}
	finalSql  string
	finalArgs []interface{}
}

func (s *SqlExpVisitor) VisitEqCond(eq *sqlcond.EqCond) error {
	s.currConds = append(s.currConds, "`"+eq.Name+"`= ?")
	s.currArgs = append(s.currArgs, eq.Value)
	return nil
}

func (s *SqlExpVisitor) VisitNeqCond(neq *sqlcond.NeqCond) error {
	s.currConds = append(s.currConds, "`"+neq.Name+"` <> ?")
	s.currArgs = append(s.currArgs, neq.Value)
	return nil
}

func (s *SqlExpVisitor) VisitRange(r *sqlcond.Range) error {
	subExps := make([]string, 0, 2)
	if r.Gt != nil {
		subExps = append(subExps, "`"+r.Name+"` > ?")
		s.currArgs = append(s.currArgs, r.Gt)
	}
	if r.Lt != nil {
		subExps = append(subExps, "`"+r.Name+"` < ?")
		s.currArgs = append(s.currArgs, r.Lt)
	}
	rexp := strings.Join(subExps, " and ")
	s.currConds = append(s.currConds, "("+rexp+")")
	return nil
}

func (s *SqlExpVisitor) VisitAndExp(and *sqlcond.AndExp) error {
	s.andSql = strings.Join(s.currConds, " and ")
	s.andArgs = s.currArgs
	s.currConds = []string{}
	s.currArgs = []interface{}{}
	return nil
}

func (s *SqlExpVisitor) VisitOrExp(or *sqlcond.OrExp) error {
	s.orSql = strings.Join(s.currConds, " or ")
	s.orArgs = s.currArgs
	s.currConds = []string{}
	s.currArgs = []interface{}{}
	return nil
}

func (s *SqlExpVisitor) VisitQuery(query *sqlcond.QueryExp) error {
	if s.andSql != "" {
		s.finalSql = s.andSql
		s.finalArgs = s.andArgs
		return nil
	}
	s.finalSql = s.orSql
	s.finalArgs = s.orArgs
	return nil
}

func (s *SqlExpVisitor) FinalSql() (sql string, args []interface{}) {
	sql = s.finalSql
	args = s.finalArgs
	return
}
