package leetcode

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
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
