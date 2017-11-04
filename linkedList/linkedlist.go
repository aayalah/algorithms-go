package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var currentNode *ListNode
	co := 0
	for l1 != nil || l2 != nil {
		sum := co
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		co = sum / 10

		newDigit := new(ListNode)
		newDigit.Val = sum

		if sum > 9 {
			newDigit.Val = sum - 10
		}

		if res == nil {
			res = newDigit
			currentNode = res
		} else {
			currentNode.Next = newDigit
			currentNode = currentNode.Next
		}
	}
	if co > 0 {
		currentNode.Next = &ListNode(1, nil)
	}
	return res
}
