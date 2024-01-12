package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		ans = ans*2 + head.Val
		head = head.Next
	}
	return ans
}

func main() {
	// Test case
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: 1}
	node1.Next = node2
	node2.Next = node3

	result := getDecimalValue(node1)
	fmt.Println(result)
}

