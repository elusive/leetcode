export interface IListNode<T> {
    val: T;
    next: IListNode<T> | null;
}

export class ListNode<T = number> implements IListNode<number> {
    constructor(public val: number, public next: IListNode<number> | null = null) { }
}

export function makeListNode(values: number[]) {
    const pointer = new ListNode(Number.MIN_SAFE_INTEGER);
    let current : ListNode<number> | null = pointer;
    for (const v of values) {
        if (current == null) continue;
        current.next = new ListNode(v);
        current = current!.next;
    }
    return pointer.next;
}
