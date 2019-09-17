package problem0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one bool
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
				one: 121,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: -121,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: 10,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: 0,
			},
			a: ans{
				one: true,
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, isPalindrome(p.one))
	}
}
