package main

import "fmt"

// From task
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Start of code to send //

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (result *ListNode) {
	result = new(ListNode)
	curr := result
	overflow := 0
	for {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		curr.Val = (overflow + val1 + val2) % 10
		overflow = (overflow + val1 + val2) / 10

		if l1 == nil && l2 == nil && overflow == 0 {
			break
		}
		curr.Next = new(ListNode)
		curr = curr.Next
	}

	return result
}

// End of code to send //

func main() {
	result := addTwoNumbers(
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
	)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Printf("\n")

	result = addTwoNumbers(
		&ListNode{0, nil},
		&ListNode{0, nil},
	)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Printf("\n")

	result = addTwoNumbers(
		&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
		&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
	)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Printf("\n")

}
