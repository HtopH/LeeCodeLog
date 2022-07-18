/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：
输入：nums = [1]
输出：[[1]]

提示：
1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/
package main

import (
	"fmt"
)

type FourFive struct {
}

func (this *FourFive) Do() {
	nums := []int{1, 2, 3}
	fmt.Println(this.permute(nums))
}

func (this *FourFive) permute(nums []int) [][]int {
	var (
		res  [][]int
		temp []int
		reDo func()
	)
	infoMap := make(map[int]int)
	for k, _ := range nums {
		infoMap[k] = 1
	}
	reDo = func() {
		if len(temp)==len(nums){
			c:=make([]int,len(temp))
			copy(c,temp)
		}
		for k, v := nums {
			if infoMap[k] == 1 {
				temp = append(temp, v)
				infoMap[k] = 0
			}
		}
	}
	reDo()
	return res
}
