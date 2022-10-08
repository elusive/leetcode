import { ListNode, makeListNode } from "./list-node";

test("Test ListNode constructor.", () => {
    const expected = new ListNode(1, new ListNode(2, null));
    expect(new ListNode(1)).not.toBe(null);
    expect(new ListNode(2).val).toBe(2);
    expect(new ListNode(1, new ListNode(2))).toEqual(expected);
});

test("Test method makeListNode returns expected value.", () => {
    const node = makeListNode([1,2,3]);
    const expected = new ListNode(1, new ListNode(2, new ListNode(3)));
    expect(makeListNode([1,2,3])).toEqual(expected);
});