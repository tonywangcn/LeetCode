package problem0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two int
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
				one: "PAYPALISHIRING",
				two: 3,
			},
			a: ans{
				one: "PAHNAPLSIIGYIR",
			},
		},
		question{
			p: para{
				one: "PAYPALISHIRING",
				two: 4,
			},
			a: ans{
				one: "PINALSIGYAHRPI",
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, convert(p.one, p.two))
	}
}
