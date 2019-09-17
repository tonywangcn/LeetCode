package problem0011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
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
				one: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			a: ans{
				one: 49,
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		// ast.Equal(a.one, maxArea(p.one))
		ast.Equal(a.one, maxAreaTwoPointer(p.one))
	}
}
