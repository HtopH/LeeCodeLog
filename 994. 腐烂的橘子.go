/*
在给定的m x n网格grid中，每个单元格可以有以下三个值之一：
值0代表空单元格；
值1代表新鲜橘子；
值2代表腐烂的橘子。
每分钟，腐烂的橘子周围4 个方向上相邻 的新鲜橘子都会腐烂。
返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回-1。

示例 1：
输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
输出：4

示例 2：
输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。

示例 3：
输入：grid = [[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 10
grid[i][j] 仅为0、1或2
*/
package main

import "fmt"

type NineNineFour struct {
}

func (this *NineNineFour) Do() {
	grid := [][]int{
		{2},
		{2},
		{1},
		{0},
		{1},
		{1}}
	fmt.Println(this.orangesRotting(grid))
}

func (this *NineNineFour) orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	have, change, num, do, lengh, wide := 0, 0, 0, 0, len(grid), len(grid[0])
	left, up, right, low := 0, 0, 0, 0
	for {
		do = 0
		//fmt.Println("----------------------------------------")
		for i := 0; i < lengh; i++ {
			//fmt.Println(grid[i])

			for j := 0; j < wide; j++ {
				left, up, right, low = 0, 0, 0, 0
				if grid[i][j] == 1 {
					have = 1
					change = 0
					if i-1 >= 0 {
						left = grid[i-1][j]
					}
					if i+1 < lengh {
						right = grid[i+1][j]
					}
					if j-1 >= 0 {
						up = grid[i][j-1]
					}
					if j+1 < wide {
						low = grid[i][j+1]
					}
					if left+up+right+low == 0 {
						return -1
					}
					//fmt.Printf("%v,%v的上下左右是%v,%v,%v,%v \n", i, j, up, low, left, right)
					if left == 2 || up == 2 || right > 1 || low > 1 {
						grid[i][j] = 3
						do = 1
						change = 1
					}
				} else if grid[i][j] == 3 {
					grid[i][j] = 2
				}
			}
		}
		if do == 1 {
			num++
		} else {
			break
		}

	}
	if have > change {
		num = -1
	}
	return num
}
