package problem0013

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
				one: "III",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "IV",
			},
			a: ans{
				one: 4,
			},
		},
		question{
			p: para{
				one: "IX",
			},
			a: ans{
				one: 9,
			},
		},
		question{
			p: para{
				one: "LVIII",
			},
			a: ans{
				one: 58,
			},
		},
		question{
			p: para{
				one: "MCMXCIV",
			},
			a: ans{
				one: 1994,
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, romanToInt(p.one))
		// ast.Equal(a.one, intToRomanDict(p.one))
	}
}
