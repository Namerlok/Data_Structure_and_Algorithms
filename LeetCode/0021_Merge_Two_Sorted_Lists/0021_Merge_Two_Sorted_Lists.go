package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Start of code to send //

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	var prev *ListNode

	for list1 != nil || list2 != nil {
		val := -101
		if list1 == nil {
			val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			val = list1.Val
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			val = list1.Val
			list1 = list1.Next
		} else {
			val = list2.Val
			list2 = list2.Next
		}

		fmt.Println(prev)
		if prev == nil {
			prev = &ListNode{val, nil}
			result = prev
		} else {
			prev.Next = &ListNode{val, nil}
			prev = prev.Next
		}
	}

	return result
}

// End of code to send //

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	result := mergeTwoLists(list1, list2)
	fmt.Printf("[")
	for result != nil {
		fmt.Printf("%d", result.Val)
		result = result.Next
		if result != nil {
			fmt.Printf(", ")
		} else {
			break
		}
	}
	fmt.Printf("]\n")
	fmt.Printf("Correct result: [1, 1, 2, 3, 4, 4]\n")

	list1 = nil
	list2 = nil
	result = mergeTwoLists(list1, list2)
	fmt.Printf("[")
	for result != nil {
		fmt.Printf("%d", result.Val)
		result = result.Next
		if result != nil {
			fmt.Printf(", ")
		} else {
			break
		}
	}
	fmt.Printf("]\n")
	fmt.Printf("Correct result: []\n")

	list1 = nil
	list2 = &ListNode{0, nil}
	result = mergeTwoLists(list1, list2)
	fmt.Printf("[")
	for result != nil {
		fmt.Printf("%d", result.Val)
		result = result.Next
		if result != nil {
			fmt.Printf(", ")
		} else {
			break
		}
	}
	fmt.Printf("]\n")
	fmt.Printf("Correct result: [0]\n")
}
