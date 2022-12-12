package leetcode

import (
	"strconv"
	"strings"
)


func newList(vs ...int) *ListNode {
	l := new(ListNode)
	n := l
	for _,v := range vs {
		n.Next = &ListNode{Val: v}
		n = n.Next
	}
	return l.Next
}


func (n *ListNode) String() string {
	if n == nil {
		return "<nil>"
	}

	var b strings.Builder

	b.WriteByte('[')
	b.WriteString(strconv.Itoa(n.Val))

	for n = n.Next; n != nil; n = n.Next {
		b.WriteString(" -> ")
		b.WriteString(strconv.Itoa(n.Val))
	}

	b.WriteByte(']')

	return b.String()
}