import { makeListNode } from "../data-structures/list-node";
import { AddTwoNumbers } from "./add-two-numbers";

test("AddTwoNumbers returns expected value.", () => {
    expect(AddTwoNumbers(makeListNode([1,2,3]), makeListNode([2,3,4]))).toEqual(makeListNode([3,5,7]));
});