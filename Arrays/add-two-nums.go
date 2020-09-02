package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//Runtime error: out of memory
func addTwoNumbers_solution1(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var current *ListNode
	var previous *ListNode

	n1, n2, carry := 0, 0, 0

	for i := 0; l1 != nil || l2 != nil || carry != 0; i++ {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		current = &ListNode{Val: (n1 + n2 + carry) % 10}
		if i == 0 {
			head = current
			previous = current
		}
		previous.Next = current
		previous = current
		carry = (n1 + n2 + carry) / 10
	}

	return head

}

func addTwoNumbers_solution2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

func main() {
	l1 := &ListNode{Val: 2}
	l2 := &ListNode{Val: 4}
	l3 := &ListNode{Val: 3}
	l1.Next = l2
	l2.Next = l3

	t1 := &ListNode{Val: 5}
	t2 := &ListNode{Val: 6}
	t3 := &ListNode{Val: 4}
	t1.Next = t2
	t2.Next = t3

	result := addTwoNumbers(l1, t1)
	for result.Next != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
	fmt.Print(result.Val)
}
