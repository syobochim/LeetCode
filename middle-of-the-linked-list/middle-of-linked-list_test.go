package middleOfLinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_middleNode_5(t *testing.T) {
	result := middleNode(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}})
	expected := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	assert.Equal(t, result, expected)
}

func Test_middleNode_6(t *testing.T) {
	result := middleNode(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}})
	expected := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}
	assert.Equal(t, result, expected)

}
