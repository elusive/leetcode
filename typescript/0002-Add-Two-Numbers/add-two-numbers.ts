import { ListNode, IListNode } from "../data-structures/list-node";

export function AddTwoNumbers(
    list1 : IListNode<number> | null, 
    list2: IListNode<number> | null
    ) : IListNode<number> | null {
        const holder = new ListNode(0);
        let rtn = holder;
        let carry : number = 0;

        while (list1 || list2 || carry) {
            if (list1) {
                carry += list1.val;
                list1 = list1.next;
            }

            if (list2) {
                carry += list2.val;
                list2 = list2.next;
            }

            if (rtn != null) {
            rtn.next = new ListNode(carry % 10);
            rtn = rtn.next;
            }
            carry = Math.floor(carry / 10);
        }

    return holder.next;
}