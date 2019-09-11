package problem0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two []int
}

type ans struct {
	one float64
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, findMedianSortedArrays(p.one, p.two))
	}
}
