/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

提示：
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
*/
package main

import (
	"fmt"
	"sort"
)

type OneNine struct {
}
type OneNineListNode struct {
	Val  int
	Next *OneNineListNode
}

func (this *OneNine) Do() {
	values := []int{1, 2}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	var node *OneNineListNode
	var temp *OneNineListNode
	for _, v := range values {
		node = &OneNineListNode{Val: v, Next: temp}
		temp = node
	}
	fmt.Println(this.RemoveNthFromEnd(node, 2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func (this *OneNine) RemoveNthFromEnd(head *OneNineListNode, n int) *OneNineListNode {
	slowNode := head
	fastNode := head
	for ; n > 0; n-- {
		fastNode = fastNode.Next
	}
	for {
		if fastNode == nil { //情况一:fastNode移动n次后为空,说明倒数n的节点为第一个节点,直接返回头部节点的下一个节点
			return slowNode.Next
		} else if fastNode.Next == nil { //情况二:fastNode移动n次后不为空,则慢指针和快指针一起移动,直至到尾部,此时慢指针的下一个节点就是要删除的节点
			slowNode.Next = slowNode.Next.Next
			break
		} else {
			fastNode = fastNode.Next
		}
		slowNode = slowNode.Next
	}
	return head
}
