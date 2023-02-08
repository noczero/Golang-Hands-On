package LeetCode

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var nbNodes int = 1
	for current := head; current.Next != nil; nbNodes++ {
		current = current.Next
	}
	middleIndex := nbNodes / 2

	var middleNode *ListNode = head
	for i := 0; i < middleIndex; i++ {
		middleNode = middleNode.Next
	}

	return middleNode
}
