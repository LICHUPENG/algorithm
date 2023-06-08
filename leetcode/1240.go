package main

//你是一位施工队的工长，根据设计师的要求准备为一套设计风格独特的房子进行室内装修。
//
// 房子的客厅大小为 n x m，为保持极简的风格，需要使用尽可能少的 正方形 瓷砖来铺盖地面。
//
// 假设正方形瓷砖的规格不限，边长都是整数。
//
// 请你帮设计师计算一下，最少需要用到多少块方形瓷砖？
//
//
//
// 示例 1：
//
//
//
// 输入：n = 2, m = 3
//输出：3
//解释：3 块地砖就可以铺满卧室。
//     2 块 1x1 地砖
//     1 块 2x2 地砖
//
// 示例 2：
//
//
//
// 输入：n = 5, m = 8
//输出：5
//
//
// 示例 3：
//
//
//
// 输入：n = 11, m = 13
//输出：6
//
//
//
//
// 提示：
//
//
// 1 <= n <= 13
// 1 <= m <= 13
//
//
// Related Topics 动态规划 回溯 👍 130 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func tilingRectangle(n int, m int) int {
	ans := max(n, m)
	rect := make([][]bool, n)
	for i := 0; i < n; i++ {
		rect[i] = make([]bool, m)
	}

	isAvailable := func(x, y, k int) bool {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				if rect[x+i][y+j] {
					return false
				}
			}
		}
		return true
	}

	fillUp := func(x, y, k int, val bool) {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				rect[x+i][y+j] = val
			}
		}
	}

	var dfs func(int, int, int)
	dfs = func(x, y, cnt int) {
		if cnt >= ans {
			return
		}
		if x >= n {
			ans = cnt
			return
		}
		// 检测下一行
		if y >= m {
			dfs(x+1, 0, cnt)
			return
		}
		// 如当前已经被覆盖，则直接尝试下一个位置
		if rect[x][y] {
			dfs(x, y+1, cnt)
			return
		}
		for k := min(n-x, m-y); k >= 1 && isAvailable(x, y, k); k-- {
			// 将长度为 k 的正方形区域标记覆盖
			fillUp(x, y, k, true)
			// 跳过 k 个位置开始检测
			dfs(x, y+k, cnt+1)
			fillUp(x, y, k, false)
		}
	}

	dfs(0, 0, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
