/*
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。

示例1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。


提示：
0 <= s.length <= 5 * 104
s由英文字母、数字、符号和空格组成
*/
package main

import (
	"fmt"
)

type Three struct {
}

func (this *Three) Do() {
	str := "pwwkew"
	fmt.Println(this.lengthOfLongestSubstring(str))
}
func (this *Three) lengthOfLongestSubstring(s string) int {
	res, left, right, end := 0, 0, 0, len(s)-1
	var temp uint8
	mark := map[uint8]int{}
	for left <= right && right <= end {
		temp = s[right]
		if mark[temp] == 0 {
			mark[temp] = right + 1
			right++
		} else {
			if mark[temp]-1 > left {
				left = mark[temp] - 1
			} else {
				left++
			}
			right = left + 1
			mark = map[uint8]int{s[left]: left + 1}
		}
		if res < right-left {
			res = right - left
		}
		if end+1-left == res {
			break
		}
	}
	return res
}
