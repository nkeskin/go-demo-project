package reverse

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO fix implementation
func reverseList(head *ListNode) *ListNode {
	end := head
	start := head
	count := 0
	for end.Next != nil {
		end = end.Next
		count++
	}
	for i := count; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			temp := start
			temp2 := start.Next
			temp.Next = temp2.Next
			temp2.Next = temp
			start = temp2
			head = start
			start = start.Next
			fmt.Println()
			//start = temp2
			//start.Val = temp2.Val
			//start.Next = temp
			//start.Next.Val = temp.Val
		}
	}
	return head
}
