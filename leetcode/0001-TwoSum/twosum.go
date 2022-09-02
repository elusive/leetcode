
package leetcode

import "fmt"

func main() {
    fmt
}

func TwoSum(nums []int, target int) [2]int {
    var sum int = 0
    for i, n := range nums {
       sum = n
       for j, x := range nums[i:] {
           if (sum + x) == target {
               return [2]int {i,j}
           }
       } 
    }
    return [2]int {-1,-1}
}
