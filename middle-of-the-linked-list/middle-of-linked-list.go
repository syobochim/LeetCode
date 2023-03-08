package middleOfLinkedList

func middleNode(head *ListNode) *ListNode {
	var result []*ListNode
	for head != nil {
		result = append(result, head)
		head = head.Next
	}
	return result[len(result)/2] // lenは整数を返し、 2で割ると小数点は切り捨てられる
}

/**
* Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
