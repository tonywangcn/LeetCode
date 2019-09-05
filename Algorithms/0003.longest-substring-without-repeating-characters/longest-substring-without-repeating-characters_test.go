package problem0003

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
				one: "abcabcbb",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "bbbbb",
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: "pwwkew",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "aab",
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: "dvdf",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "",
			},
			a: ans{
				one: 0,
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, lengthOfLongestSubstring_BruteForce(p.one))
		ast.Equal(a.one, lengthOfLongestSubstring_HashMap(p.one))
	}
}
