/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
示例 1：
输入：nums = [-4,-1,0,10,3]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]

提示：
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 已按 非递减顺序 排序
*/
package main

import (
	"fmt"
)

type NineSevenSeven struct {
}

func (this *NineSevenSeven) Do() {
	nums := []int{-1, 2, -3, 4, -5, -6, 7, -8, 9, -10, 11}
	fmt.Println(this.SortedSquares(nums))
}

func (this *NineSevenSeven) SortedSquares(nums []int) []int {
	for k, v := range nums {
		nums[k] = v * v
	}
	this.QuickSort(nums, 0, len(nums)-1)
	return nums
}

/*
快速排序算法通过多次比较和交换来实现排序，其排序流程如下： [2]
(1)首先设定一个分界值，通过该分界值将数组分成左右两部分。 [2]
(2)将大于或等于分界值的数据集中到数组右边，小于分界值的数据集中到数组的左边。此时，左边部分中各元素都小于分界值，而右边部分中各元素都大于或等于分界值。 [2]
(3)然后，左边和右边的数据可以独立排序。对于左侧的数组数据，又可以取一个分界值，将该部分数据分成左右两部分，同样在左边放置较小值，右边放置较大值。右侧的数组数据也可以做类似处理。 [2]
(4)重复上述过程，可以看出，这是一个递归定义。通过递归将左侧部分排好序后，再递归排好右侧部分的顺序。当左、右两个部分各数据排序完成后，整个数组的排序也就完成了。
*/

func (this *NineSevenSeven) QuickSort(arr []int, left int, right int) {
	if left < right {
		p := left
		temp := arr[p]
		high, low := right, left
		fmt.Println("start:", arr, p)
		for low < high {
			fmt.Println(arr, low, high, p)
			for high > p && arr[high] > temp {
				fmt.Println("high--")
				high--
			}
			if high > p {
				fmt.Println("high change start:", arr, high, p)
				arr[p], arr[high] = arr[high], arr[p]
				p = high
				fmt.Println("high change end:", arr, high, p)
			}
			for low < p && arr[low] < temp {
				fmt.Println("low++")
				low++
			}
			if low < p {
				fmt.Println("low change start:", arr, low, p)
				arr[p], arr[low] = arr[low], arr[p]
				p = low
				fmt.Println("low change end:", arr, low, p)
			}
		}
		fmt.Println("end:", arr, p)
		if p > left {
			this.QuickSort(arr, left, p)
		}
		if p+1 < right {
			this.QuickSort(arr, p+1, right)
		}
	}

}
