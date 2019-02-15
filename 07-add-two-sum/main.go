/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	head := cur

	sum, co := 0, 0
	for co != 0 || l1 != nil || l2 != nil {
		cur.Next = &ListNode{}
		cur = cur.Next

		sum = co
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		co = sum / 10
		cur.Val = sum % 10
	}
	return head.Next
}