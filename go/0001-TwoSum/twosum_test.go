
package leetcode

import "testing"

type Test struct{
    description    string
    nums           []int
    target         int
    expected       [2]int
}

/*
  Constraints:
    2 <= nums.length <= 10^4
    -10^9 <= nums[i] <= 10^9
    -10^9 <= target <= 10^9
*/


func getExamples() []Test {
    var tests = []Test {
        {"Example 1", []int{2,7,11,15}, 9, [2]int{0,1} },
        {"Example 2", []int{3,2,4}, 6, [2]int{1,2}},
        {"Example 3", []int{3,3}, 6, [2]int{0,1}},
        {"Test 1", []int{8,12,4,3,6}, 7, [2]int{2,3}},
    }
    return tests
}


func TestExamples(t *testing.T) {
    tests := getExamples()
    for _, tv := range tests {
        t.Run(tv.description, func(t *testing.T) {
            result := TwoSum(tv.nums, tv.target)
            if result != tv.expected {
                t.Errorf("Wrong result! Expected: %d, returned: %d ", tv.expected, result)
            }
        })
    }
}
