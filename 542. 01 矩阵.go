/*
给定一个由 0 和 1 组成的矩阵 mat，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
两个相邻元素间的距离为 1 。

示例 1：
输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
输出：[[0,0,0],[0,1,0],[0,0,0]]

示例 2：
输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
输出：[[0,0,0],[0,1,0],[1,2,1]]


提示：
m == mat.length
n == mat[i].length
1 <= m, n <= 104
1 <= m * n <= 104
mat[i][j] is either 0 or 1.
mat 中至少有一个 0
*/

package main

import "fmt"

type FiveFourTwo struct {
}

func (this *FiveFourTwo) Do() {
	mat := [][]int{
		{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 1, 0},
		{1, 1, 1, 1, 0, 1, 0, 0, 1, 1}}
	fmt.Println(this.updateMatrix(mat))
}
func (this *FiveFourTwo) updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}
	var temp int
	var mark int
	total := len(mat) * len(mat[0])
	for mark < total {
		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[0]); j++ {
				if mat[i][j] == 0 && temp == 0 {
					mark++
					continue
				}
				if mat[i][j] != temp {
					continue
				}
				if (i-1 >= 0 && mat[i-1][j] == temp-1) || (i+1 < len(mat) && mat[i+1][j] == temp-1) || (j-1 >= 0 && mat[i][j-1] == temp-1) || (j+1 < len(mat[0]) && mat[i][j+1] == temp-1) {
					mark++
				} else {
					mat[i][j] = temp + 1
				}

			}
		}
		temp++
	}

	return mat
}
