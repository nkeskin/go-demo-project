package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listNode := ListNode{4, nil}
	listNode.Next = &ListNode{5, nil}
	listNode.Next.Next = &ListNode{1, nil}
	listNode.Next.Next.Next = &ListNode{9, nil}
	deletedNode := listNode.Next
	deleteNode(deletedNode)
	for listNode != (ListNode{}) {
		fmt.Println(listNode.Val)
		listNode = *listNode.Next
	}
}

// Dereference the node to reflect changes on original variable
func deleteNode(node *ListNode) {
	*node = *node.Next
}
