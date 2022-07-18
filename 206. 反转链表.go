/*

 */
package main

import "fmt"

type TwoZeroSix struct {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *TwoZeroSix) Do() {
	one := []int{1, 2, 3}
	var list1 *ListNode
	for _, v := range one {
		list1 = &ListNode{
			Val:  v,
			Next: list1,
		}
	}
	res := this.reverseList(list1)
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
func (this *TwoZeroSix) reverseList(head *ListNode) *ListNode {
	var temp *ListNode
	for head != nil {
		//fmt.Println(head, temp)
		temp = &ListNode{
			Val:  head.Val,
			Next: temp,
		}
		head = head.Next
	}
	return temp
}
