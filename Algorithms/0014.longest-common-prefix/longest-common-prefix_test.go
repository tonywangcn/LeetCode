package problem0014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []string
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
				one: []string{"flower", "flow", "flight"},
			},
			a: ans{
				one: "fl",
			},
		},
		question{
			p: para{
				one: []string{"dog", "racecar", "car"},
			},
			a: ans{
				one: "",
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, longestCommonPrefix(p.one))
	}
}
