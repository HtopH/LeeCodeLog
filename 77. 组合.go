/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。
示例 1：
输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

示例 2：
输入：n = 1, k = 1
输出：[[1]]
提示：
1 <= n <= 20
1 <= k <= n
*/
package main

import "fmt"

type SevenSeven struct {
}

func (this *SevenSeven) Do() {
	fmt.Println(this.combine(5, 3))
}

func (this *SevenSeven) combine(n int, k int) [][]int {
	var (
		res  [][]int
		temp []int
	)
	var reDo func(l int)
	reDo = func(l int) {
		if len(temp)+(n-l+1) < k {
			return
		}
		if len(temp) == k {
			c := make([]int, k)
			copy(c, temp)
			res = append(res, c)
			return
		}
		temp = append(temp, l)
		reDo(l + 1)
		temp = temp[:len(temp)-1]
		reDo(l + 1)

	}
	reDo(1)
	return res
}
