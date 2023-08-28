package leetcode

import (
	"fmt"
	"testing"
)


/*
   Constraints:
       linked list length [1,100]
       0 <= node.Val <= 9
       no leading zeros in represented numbers
*/


func TestExamples(t *testing.T) {
	var test = []struct{
        description     string
        input1, input2  *ListNode
        expected        *ListNode
    }{
        {"Test Case 1", NewList(0, 1), NewList(0, 1, 2), NewList(0, 2, 2)},
        /*{"Test Case 2", &ListNode{}, NewList(0, 1), NewList(0, 1)},*/
        {"Test Case 3", NewList(9, 9), NewList(1), NewList(0, 0, 1)},
        {"Example 1", NewList(2, 4, 3), NewList(5, 6, 4), NewList(7, 0, 8)},
        {"Example 2", NewList(0), NewList(0), NewList(0)},
        {"Example 3", NewList(9, 9, 9, 9, 9, 9, 9), NewList(9, 9, 9, 9), NewList(8, 9, 9, 9, 0, 0, 0, 1)},
    }
	fmt.Printf("%v\n", test)
	for _, tv := range test {
		t.Run(tv.description, func(t *testing.T) {
			result := AddTwoNumbers(tv.input1, tv.input2)
			if result != tv.expected {
				t.Errorf("Wrong result! Expected: %+v, returned: %+v ", tv.expected, result)
			}
		})
	}
}
