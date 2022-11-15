package add_two_numbers

import (
	"fmt"
	"testing"
)

type Test struct {
	description string
	input1      ListNode
	input2      ListNode
	expected    ListNode
}

/*
   Constraints:
       linked list length [1,100]
       0 <= node.Val <= 9
       no leading zeros in represented numbers
*/

func getExamples() []Test {
	var tests = []Test{
		{
			"Test Case 1",
			ListNode{Val: 0, Next: &ListNode{Val: 1}},
			ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}},
		},
		{
			"Test Case 2",
			ListNode{},
			ListNode{Val: 0, Next: &ListNode{Val: 1}},
			ListNode{Val: 0, Next: &ListNode{Val: 1}},
		},
		{
			"Test Case 3",
			ListNode{Val: 9, Next: &ListNode{Val: 9}},
			ListNode{Val: 1},
			ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			"Example 1",
			ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			"Example 2",
			ListNode{Val: 0},
			ListNode{Val: 0},
			ListNode{Val: 0},
		},
		{
			"Example 3",
			ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}},
			ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
			ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}},
		},
	}
	return tests
}

func TestExamples(t *testing.T) {
	tests := getExamples()
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
