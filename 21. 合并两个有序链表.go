/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]

提示：
两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/

package main

import (
	"fmt"
	"sort"
)

type TwoOne struct {
}
type TwoOneListNode struct {
	Val  int
	Next *TwoOneListNode
}

func (this *TwoOne) Do() {
	one := []int{1, 2, 3}
	two := []int{2}
	sort.Slice(one, func(i, j int) bool {
		return one[i] > one[j]
	})
	sort.Slice(two, func(i, j int) bool {
		return two[i] > two[j]
	})
	var list1 *TwoOneListNode
	var list2 *TwoOneListNode
	for _, v := range one {
		list1 = &TwoOneListNode{
			Val:  v,
			Next: list1,
		}
	}
	for _, v := range two {
		list2 = &TwoOneListNode{
			Val:  v,
			Next: list2,
		}
	}
	res := this.mergeTwoLists(list1, list2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func (this *TwoOne) mergeTwoLists(list1 *TwoOneListNode, list2 *TwoOneListNode) *TwoOneListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list2.Val > list1.Val {
		list1, list2 = list2, list1
	}
	var head = list2
	for list1 != nil {

		//fmt.Println("-------------------------------")
		for list2 != nil && list2.Val <= list1.Val {
			//fmt.Println("list2 ", list2.Val, list2.Next)
			if list2.Next == nil {
				//fmt.Println("list2 分支1", list2.Val, list2.Next)
				list2.Next = list1
				return head
			} else if list2.Next.Val < list1.Val {
				list2 = list2.Next
				//fmt.Println("list2 分支2", list2.Val, list2.Next)
			} else {
				list2.Next = &TwoOneListNode{
					Val:  list1.Val,
					Next: list2.Next,
				}
				//fmt.Println("list2 分支3", list2.Val, list2.Next, list2.Next.Next)
				list2 = list2.Next
				break
			}
		}
		list1 = list1.Next
	}
	return head
}
