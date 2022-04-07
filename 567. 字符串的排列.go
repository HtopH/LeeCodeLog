/*
给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").

示例 2：
输入：s1= "ab" s2 = "eidboaoo"
输出：false


提示：
1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母
*/
package main

import (
	"fmt"
)

type FineSixSeven struct {
}

func (this *FineSixSeven) Do() {
	s1 := "ab"
	s2 := "a"
	fmt.Println(this.CheckInclusion(s1, s2))
}

func (this *FineSixSeven) CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	mark := map[int32]int{}
	doMark := make([]int, len(s2), len(s2))
	nums := 0
	for _, v := range s1 {
		if mark[v] == 0 {
			mark[v] = 2
		} else {
			mark[v] += 1
		}
		nums++
	}
	//fmt.Println("start:", mark, nums)
	left, right := 0, len(s1)-1
	for {
		for i := left; i <= right; i++ {
			if doMark[i] == 1 {
				continue
			}
			if mark[int32(s2[i])] > 1 {
				mark[int32(s2[i])] -= 1
				nums--
				doMark[i] = 1
			}
		}
		//fmt.Println(left, right, mark, nums, doMark)
		if nums == 0 {
			return true
		} else {
			left++
			right++
			if right == len(s2) {
				break
			}

		}
		if doMark[left-1] == 1 {
			doMark[left-1] = 0
			mark[int32(s2[left-1])] += 1
			nums++
		}

	}
	return false
}
