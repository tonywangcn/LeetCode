package problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one *ListNode
	two *ListNode
}

type ans struct {
	one *ListNode
}

type question struct {
	p para
	a ans
}

func makeListNode(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}
	res := &ListNode{
		Val: data[0],
	}
	tmp := res
	for i := 1; i < len(data); i++ {
		tmp.Next = &ListNode{Val: data[i]}
		tmp = tmp.Next
	}
	return res

}

func Test_OK(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{
			p: para{
				one: makeListNode([]int{2, 4, 3}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: makeListNode([]int{7, 0, 8}),
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, addTwoNumbers(p.one, p.two))
	}
}
