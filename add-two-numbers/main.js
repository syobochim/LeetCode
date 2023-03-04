function addTwoNumbers(l1, l2) {
    var dummyHead = new ListNode(0);
    var currentHead = dummyHead;
    var carry = 0;
    while (l1 !== null || l2 !== null || carry !== 0) {
        var x = l1 !== null ? l1.val : 0;
        var y = l2 !== null ? l2.val : 0;
        var sum = x + y + carry;
        carry = Math.floor(sum / 10);
        currentHead.next = new ListNode(sum % 10);
        currentHead = currentHead.next;
        if (l1 !== null) {
            l1 = l1.next;
        }
        if (l2 !== null) {
            l2 = l2.next;
        }
    }
    return dummyHead.next;
}
// Definition for singly-linked list.
var ListNode = /** @class */ (function () {
    function ListNode(val, next) {
        this.val = (val === undefined ? 0 : val);
        this.next = (next === undefined ? null : next);
    }
    return ListNode;
}());
function test() {
    console.log(addTwoNumbers(new ListNode(2, new ListNode(4, new ListNode(3))), new ListNode(5, new ListNode(6, new ListNode(4)))));
    console.log(addTwoNumbers(new ListNode(0), new ListNode(0)));
    console.log(addTwoNumbers(new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9))))))), new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9))))));
}
test();
