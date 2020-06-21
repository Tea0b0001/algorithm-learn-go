package common

import "strconv"

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func List2Str(head * ListNode) string {
	var listStr string = ""
	cur := head

	for ;cur != nil; {
		listStr += strconv.Itoa(cur.Val)
		if cur.Next != nil {
			listStr += "->"
		}
		cur = cur.Next
	}
	return listStr
}