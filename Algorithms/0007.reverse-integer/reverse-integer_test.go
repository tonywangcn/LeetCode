package problem0007

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
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
				one: 123,
			},
			a: ans{
				one: 321,
			},
		},
		question{
			p: para{
				one: -123,
			},
			a: ans{
				one: -321,
			},
		},
		question{
			p: para{
				one: 120,
			},
			a: ans{
				one: 21,
			},
		},
		question{
			p: para{
				one: 1534236469,
			},
			a: ans{
				one: 0,
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, reverse(p.one))
	}
}
