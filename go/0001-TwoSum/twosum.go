
package leetcode

func TwoSum(nums []int, target int) [2]int {
    var sum int = 0
    for i, n := range nums {
       sum = n
       for j, x := range nums[i+1:] {
           if (sum + x) == target {
                return [2]int {i,i+j+1}
           }
       } 
    }
    return [2]int {-1,-1}
}
