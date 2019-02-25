package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ListNode{
		1,
		nil,
	}
	b := ListNode{
		2,
		nil,
	}

	a.Next = &b

	c := ListNode{
		3,
		nil,
	}
	b.Next = &c

	head := ListNode{
		Next: &a,
		Val:  1,
	}
	r := reverse(&head)
	for {
		fmt.Println(r.Val)
		r = r.Next

		if r == nil {
			return
		}
	}

}
func reverse(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	ret := head
	head = head.Next
	ret.Next = nil
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = ret

		ret = tmp
	}
	return ret
}
