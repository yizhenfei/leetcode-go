package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	result := ListNode{}
	var result *ListNode

	for l1 != nil || l2 != nil || carry != 0 {
		// Add current digits and carry
		current := carry
		if l1 != nil {
			current += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			current += l2.Val
			l2 = l2.Next
		}

		// Process carry
		carry = 0
		if current >= 10 {
			current -= 10
			carry = 1
		}

		// Append result
		result = &ListNode{
			Val:  current,
			Next: result,
		}
	}

	return result
}
