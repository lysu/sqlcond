package sql

import (
	"encoding/json"
	"fmt"
	"testing"
	"github.com/lysu/sqlcond"
)

func TestName(t *testing.T) {

	x := `{
        "and": [{
            "eq": {
                "name": "nick",
                "value": "robi"
            }
        }, {
            "neq": {
                "name": "age",
                "value": 2333
            }
        }, {
            "range": {
                "name": "age_leve",
                "gt": 1,
                "lt": 3
            }
        }]
    }`
	fmt.Println(x)
	var x2 sqlcond.QueryExp
	err := json.Unmarshal([]byte(x), &x2)
	if err != nil {
		t.Error(err)
	}
	v := &SqlExpVisitor{}
	x2.Accept(v)
	fmt.Println("================>")
	fmt.Println(v.FinalSql())

}
