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
        {"Test Case 1", newList(0, 1), newList(0, 1, 2), newList(0, 2, 2)},
        /*{"Test Case 2", &ListNode{}, newList(0, 1), newList(0, 1)},*/
        {"Test Case 3", newList(9, 9), newList(1), newList(0, 0, 1)},
        {"Example 1", newList(2, 4, 3), newList(5, 6, 4), newList(7, 0, 8)},
        {"Example 2", newList(0), newList(0), newList(0)},
        {"Example 3", newList(9, 9, 9, 9, 9, 9, 9), newList(9, 9, 9, 9), newList(8, 9, 9, 9, 0, 0, 0, 1)},
    }
	fmt.Printf("%v\n", tests)
	for _, tv := range tests {
		t.Run(tv.description, func(t *testing.T) {
			result := AddTwoNumbers(tv.input1, tv.input2)
			if result != tv.expected {
				t.Errorf("Wrong result! Expected: %+v, returned: %+v ", tv.expected, result)
			}
		})
	}
}
