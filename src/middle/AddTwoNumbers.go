package main

import "common"

/*
	@title:两数相加
	@description:给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	@example:
	示例：
	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/add-two-numbers
 */

/*
	@version:Go 1.13.12
*/
func main()  {
	l1First := common.ListNode{
		Val: 2,
	}
	l1Second := common.ListNode{
		Val: 4,
	}
	l1Third := common.ListNode{
		Val: 3,
	}
	l1First.Next = &l1Second
	l1Second.Next = &l1Third

	l2First := common.ListNode{
		Val: 5,
	}
	l2Second := common.ListNode{
		Val: 6,
	}
	l2Third := common.ListNode{
		Val: 4,
	}
	l2First.Next = &l2Second
	l2Second.Next = &l2Third

	result := addTwoNumbers(&l1First, &l2First)
	print(common.List2Str(&l1First), " + ", common.List2Str(&l2First))
	print(" = ", common.List2Str(result))
}

/*
	就像我们平时在做加法一样，往往从两个数字的最低位开始相加，得到两个位数的加法结果和进位情况，然后把进位带到下一位数的加法中，以此类推。
	我们可以从两个列表的头部开始，做位数相加，并且要记录进位值。同时要考虑两个数字位数不同的问题，这里使用用0值代替空位的方法来解决位数不同的问题。
 */
func addTwoNumbers(l1 * common.ListNode, l2 * common.ListNode) * common.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	resultHead := common.ListNode{
		Val: 0,
		Next: nil,
	}

	var cur * common.ListNode
	cur = &resultHead
	var carry int = 0

	for ; l1 != nil || l2 != nil || carry == 1; {
		//这里是为了解决数字位数不同的问题，相当于短位数做了补0的操作
		var l1Val int = 0
		var l2Val int = 0
		if l1 != nil {
			l1Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val += l2.Val
			l2 = l2.Next
		}
		
		sumNode := l1Val + l2Val + carry
		carry = sumNode / 10

		newListNode := common.ListNode{
			Val: sumNode % 10,
			Next: nil,
		}

		cur.Next = &newListNode
		cur = cur.Next

	}
	return resultHead.Next
}