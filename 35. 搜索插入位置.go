/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。

示例 1:
输入: nums = [1,3,5,6,8,11,16], target = 5
输出: 2

示例2:
输入: nums = [1,3,5,6], target = 2
输出: 1

示例 3:
输入: nums = [1,3,5,6], target = 7
输出: 4

提示:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 为无重复元素的升序排列数组
-104 <= target <= 104

*/
package main

import (
	"fmt"
)

type threeFive struct {
}

func (this *threeFive) Do() {
	num := []int{1, 3, 5, 6}
	target := 1
	res := this.SearchInsertNew(num, target)
	fmt.Println(res)
}

func (this *threeFive) SearchInsert(nums []int, target int) int {
	var mid int
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid = (high-low)/2 + low
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] && mid == 0 {
			return mid
		} else if target > nums[mid] && mid == len(nums)-1 {
			return len(nums)
		} else if target > nums[mid] && target <= nums[mid+1] {
			return mid + 1
		} else if target < nums[mid] && target > nums[mid-1] {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

/*
思路与算法

假设题意是叫你在排序数组中寻找是否存在一个目标值，那么训练有素的读者肯定立马就能想到利用二分法在 O(\log n)O(logn) 的时间内找到是否存在目标值。但这题还多了个额外的条件，即如果不存在数组中的时候需要返回按顺序插入的位置，那我们还能
用二分法么？答案是可以的，我们只需要稍作修改即可。
考虑这个插入的位置 {pos}pos，它成立的条件为：{nums}[pos-1]<{target}≤nums[pos]
其中 {nums}nums 代表排序数组。由于如果存在这个目标值，我们返回的索引也是 {pos}pos，因此我们可以将两个条件合并得出最后的目标：「在一个有序数组中找第一个大于等于 {target}target 的下标」。

*/
func (this *threeFive) SearchInsertNew(nums []int, target int) int {
	var mid int
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid = (high-low)>>1 + low
		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return mid
}
