package reverse

import (
	"testing"
)

func Test_reverseList(t *testing.T) {

	linkedList := ListNode{
		Val:  5,
		Next: nil,
	}

	linkedList.Next = &ListNode{
		10, nil,
	}

	linkedList.Next.Next = &ListNode{
		15, nil,
	}

	reverseList(&linkedList)

	//type args struct {
	//	head *ListNode
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want *ListNode
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("reverseList() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
