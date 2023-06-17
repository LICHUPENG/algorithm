package main

//二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
//
// 请返回 封闭岛屿 的数目。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,
//0,1],[1,1,1,1,1,1,1,0]]
//输出：2
//解释：
//灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
//
// 示例 2：
//
//
//
//
//输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
//输出：1
//
//
// 示例 3：
//
//
//输入：grid = [[1,1,1,1,1,1,1],
//             [1,0,0,0,0,0,1],
//             [1,0,1,1,1,0,1],
//             [1,0,1,0,1,0,1],
//             [1,0,1,1,1,0,1],
//             [1,0,0,0,0,0,1],
//             [1,1,1,1,1,1,1]]
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= grid.length, grid[0].length <= 100
// 0 <= grid[i][j] <=1
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 196 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				qu := [][]int{{i, j}}
				grid[i][j] = 1
				closed := true

				for len(qu) > 0 {
					cx, cy := qu[0][0], qu[0][1]
					qu = qu[1:]

					if cx == 0 || cy == 0 || cx == m-1 || cy == n-1 {
						closed = false
					}

					directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
					for _, dir := range directions {
						nx, ny := cx+dir[0], cy+dir[1]
						if 0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == 0 {
							grid[nx][ny] = 1
							qu = append(qu, []int{nx, ny})
						}
					}
				}

				if closed {
					ans++
				}
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
