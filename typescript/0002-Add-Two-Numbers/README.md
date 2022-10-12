# LeetCode Problems
Leetcode practice problems readme and/or instructions for the problem solved in this directory. 
Link to this problem on [leetcode.com](https://leetcode.com/problems/two-sum/).

## Problem 0002 - Add Two Numbers
Leetcode problem #1 is called "Add Two Numbers" and involves adding together two linked lists of numbers. 

### Instruction Text
Given an array of integers nums and an integer target, return indices of the two numbers such 
that they add up to target. You may assume that each input would have exactly one solution, and 
you may not use the same element twice. You can return the answer in any order.

NOTE: Assumes existence of a ListNode type or interface with similar structure to this:

```typescript
export interface IListNode<T> {
    val: T;
    next: IListNode<T> | null;
}
```

### Constraints
- The number of nodes in each linked list is in the range [1, 100].
- `0 <= Node.val <= 9`
- It is guaranteed that the list represents a number that does not have leading zeros.

### Examples
The following examples are also used in the unit tests.

##### Example 1
> Input: nums = [2,7,11,15], target = 9
> Output: [0,1]
> Explanation: Because nums[0] + nums[1] == 9, we return [0, 1]. 

##### Example 2:
> Input: l1 = [0], l2 = [0]
> Output: [0]

##### Example 3
> Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
> Output: [8,9,9,9,0,0,0,1]
