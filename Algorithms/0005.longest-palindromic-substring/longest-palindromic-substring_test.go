package problem0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
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
				one: "babad",
			},
			a: ans{
				one: "bab",
			},
		},
		question{
			p: para{
				one: "cbbd",
			},
			a: ans{
				one: "bb",
			},
		},
		question{
			p: para{
				one: "bb",
			},
			a: ans{
				one: "bb",
			},
		},
		question{
			p: para{
				one: "a",
			},
			a: ans{
				one: "a",
			},
		},
		question{
			p: para{
				one: "",
			},
			a: ans{
				one: "",
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		// ast.Equal(a.one, longestPalindrome(p.one))
		ast.Equal(a.one, longestPalindromeDP(p.one))
	}
}
