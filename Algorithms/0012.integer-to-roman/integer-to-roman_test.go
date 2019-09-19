package problem0012

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one string
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
				one: 3,
			},
			a: ans{
				one: "III",
			},
		},
		question{
			p: para{
				one: 4,
			},
			a: ans{
				one: "IV",
			},
		},
		question{
			p: para{
				one: 9,
			},
			a: ans{
				one: "IX",
			},
		},
		question{
			p: para{
				one: 58,
			},
			a: ans{
				one: "LVIII",
			},
		},
		question{
			p: para{
				one: 1994,
			},
			a: ans{
				one: "MCMXCIV",
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, intToRoman(p.one))
		ast.Equal(a.one, intToRomanDict(p.one))
	}
}
