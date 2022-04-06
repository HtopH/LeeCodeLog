/*
给你一个数组，将数组中的元素向右轮转 k个位置，其中k是非负数。
示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

提示：
1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105
*/
package main

const tt int64 = 3

type OneEightNine struct {
}

func (this *OneEightNine) Do() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 5
	this.rotate(nums, k)
}

func (this *OneEightNine) rotate(nums []int, k int) {
	k = k % len(nums)
	var temp []int

	//解法一
	//for key, value := range temp {
	//	if key+k <= len(temp)-1 {
	//		nums[key+k] = value
	//	} else {
	//		nums[key+k-(len(temp))] = value
	//	}
	//}

	//解法二
	temp = append(temp, nums[len(nums)-k:]...)
	temp = append(temp, nums[:len(nums)-k]...)
	nums = temp
}
