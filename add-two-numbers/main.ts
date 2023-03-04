function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
    let dummyHead: ListNode = new ListNode(0)
    let currentHead: ListNode = dummyHead
    let carry: number = 0

    while (l1 !== null || l2 !== null || carry !== 0) {
        let x: number = l1 !== null ? l1.val : 0
        let y: number = l2 !== null ? l2.val : 0
        let sum: number = x + y + carry
        carry = Math.floor(sum / 10)
        currentHead.next = new ListNode(sum % 10)
        currentHead = currentHead.next
        if (l1 !== null) {
            l1 = l1.next
        }
        if (l2 !== null) {
            l2 = l2.next
        }
    }

    return dummyHead.next
}

// Definition for singly-linked list.
class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}

function test(): void {
    console.log(addTwoNumbers(new ListNode(2, new ListNode(4, new ListNode(3))),
                              new ListNode(5, new ListNode(6, new ListNode(4)))))

    console.log(addTwoNumbers(new ListNode(0),
                              new ListNode(0)))

    console.log(
            addTwoNumbers(
            new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9))))))),
                              new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9))))))

}
test();