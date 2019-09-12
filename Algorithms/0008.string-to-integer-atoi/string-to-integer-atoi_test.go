package problem0008

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{
			p: para{
				one: "42",
			},
			a: ans{
				one: 42,
			},
		},
		question{
			p: para{
				one: "   -42",
			},
			a: ans{
				one: -42,
			},
		},
		question{
			p: para{
				one: "4193 with words",
			},
			a: ans{
				one: 4193,
			},
		},
		question{
			p: para{
				one: "words and 987",
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: "-91283472332",
			},
			a: ans{
				one: -2147483648,
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, myAtoi(p.one))
	}
}
