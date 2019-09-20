package problem0015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one [][]int
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
				one: []int{-1, 0, 1, 2, -1, -4},
			},
			a: ans{
				one: [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}},
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, threeSum(p.one))
	}
}
