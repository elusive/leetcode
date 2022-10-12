
package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func AddTwoNumbers(input1 *ListNode, input2 *ListNode) *ListNode {
    var rtn *ListNode = &ListNode{ Val: 0 }
    var x, y, carry, sum int = 0, 0, 0, 0
    
    node := rtn
    ln1 := input1
    ln2 := input2

    for ln1 != nil || ln2 != nil || carry != 0 {
        if (ln1 != nil) { 
            x = ln1.Val 
            ln1 = ln1.Next
        }
        if (ln2 != nil) { 
            y = ln2.Val 
            ln2 = ln2.Next
        }

        sum = x + y + carry
        carry = sum / 10
        node.Next = &ListNode { Val: sum /10 } 
        
        node = node.Next 
        if (ln1 != nil) {ln1 = ln1.Next}
        if (ln2 != nil) {ln2 = ln2.Next}
    }

    return rtn
}
