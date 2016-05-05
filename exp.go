package sqlcond

type Exp interface {
	Accept(ExpVisitor) error
}

type QueryExp struct {
	And AndExp `json:"and,omitempty"`
	Or  OrExp  `json:"or,omitempty"`
}

func (q *QueryExp) Accept(v ExpVisitor) error {
	if q.And != nil {
		err := q.And.Accept(v)
		if err != nil {
			return err
		}
	}
	if q.Or != nil {
		err := q.Or.Accept(v)
		if err != nil {
			return err
		}
	}
	return v.VisitQuery(q)
}

type AndExp []Cond

func (a *AndExp) Accept(v ExpVisitor) error {
	for _, exp := range *a {
		err := exp.Accept(v)
		if err != nil {
			return err
		}
	}
	return v.VisitAndExp(a)
}

type OrExp []Cond

func (o *OrExp) Accept(v ExpVisitor) error {
	for _, exp := range *o {
		err := exp.Accept(v)
		if err != nil {
			return err
		}
	}
	return v.VisitOrExp(o)
}

type Cond struct {
	Eq    *EqCond  `json:"eq,omitempty"`
	Neq   *NeqCond `json:"neq,omitempty"`
	Range *Range   `json:"range,omitempty"`
}

func (c *Cond) Accept(v ExpVisitor) error {
	if c.Eq != nil {
		return c.Eq.Accept(v)
	}
	if c.Neq != nil {
		return c.Neq.Accept(v)
	}
	if c.Range != nil {
		return c.Range.Accept(v)
	}
	return nil
}

type EqCond struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func (e *EqCond) Accept(v ExpVisitor) error {
	return v.VisitEqCond(e)
}

type NeqCond struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func (n *NeqCond) Accept(v ExpVisitor) error {
	return v.VisitNeqCond(n)
}

type Range struct {
	Name string      `json:"name"`
	Gt   interface{} `json:"gt, omitmepty"`
	Lt   interface{} `json:"lt, omitmepty"`
}

func (r *Range) Accept(v ExpVisitor) error {
	return v.VisitRange(r)
}

type ExpVisitor interface {
	VisitEqCond(eq *EqCond) error
	VisitNeqCond(neq *NeqCond) error
	VisitRange(r *Range) error
	VisitAndExp(and *AndExp) error
	VisitOrExp(or *OrExp) error
	VisitQuery(query *QueryExp) error
}
