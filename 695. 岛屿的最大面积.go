/*
给你一个大小为 m x n 的二进制矩阵 grid 。
岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设grid 的四个边缘都被 0（代表水）包围着。
岛屿的面积是岛上值为 1 的单元格的数目。
计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

示例 1：
输入：grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
输出：6
解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。

示例 2：
输入：grid = [[0,0,0,0,0,0,0,0]]
输出：0

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 50
grid[i][j] 为 0 或 1
*/
package main

import (
	"fmt"
)

type SixNineFive struct {
}

func (this *SixNineFive) Do() {
	image := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(this.maxAreaOfIsland(image))
}

func (this *SixNineFive) maxAreaOfIsland(grid [][]int) int {
	temp := 0
	res := 0
	do := 2
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				temp = 0
				temp += sixNineFiveStat(grid, i, j, do)
				if temp > res {
					res = temp
				}
			}
		}
	}
	return res
}

func sixNineFiveStat(grid [][]int, x, y, do int) int {
	res := 0
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != 1 {
		return res
	}
	res += 1
	grid[x][y] = do
	res += sixNineFiveStat(grid, x-1, y, do)
	res += sixNineFiveStat(grid, x+1, y, do)
	res += sixNineFiveStat(grid, x, y-1, do)
	res += sixNineFiveStat(grid, x, y+1, do)
	return res
}
