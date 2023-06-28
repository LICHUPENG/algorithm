package main

import "sort"

//给你一个整数数组 nums，请你找出并返回能被三整除的元素最大和。
//
//
//
//
//
//
// 示例 1：
//
// 输入：nums = [3,6,5,1,8]
//输出：18
//解释：选出数字 3, 6, 1 和 8，它们的和是 18（可被 3 整除的最大和）。
//
// 示例 2：
//
// 输入：nums = [4]
//输出：0
//解释：4 不能被 3 整除，所以无法选出数字，返回 0。
//
//
// 示例 3：
//
// 输入：nums = [1,2,3,4,4]
//输出：12
//解释：选出数字 1, 3, 4 以及 4，它们的和是 12（可被 3 整除的最大和）。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 4 * 10^4
// 1 <= nums[i] <= 10^4
//
//
// Related Topics 贪心 数组 动态规划 排序 👍 244 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func accumulate(v []int) int {
	ans := 0
	for _, x := range v {
		ans += x
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSumDivThree(nums []int) int {
	// 使用 v[0], v[1], v[2] 分别表示 a, b, c
	v := make([][]int, 3)
	for _, num := range nums {
		v[num%3] = append(v[num%3], num)
	}
	sort.Slice(v[1], func(i, j int) bool {
		return v[1][i] > v[1][j]
	})
	sort.Slice(v[2], func(i, j int) bool {
		return v[2][i] > v[2][j]
	})
	ans, lb, lc := 0, len(v[1]), len(v[2])
	for cntb := max(lb-2, 0); cntb <= lb; cntb++ {
		for cntc := max(lc-2, 0); cntc <= lc; cntc++ {
			if (cntb-cntc)%3 == 0 {
				ans = max(ans, accumulate(v[1][:cntb])+accumulate(v[2][:cntc]))
			}
		}
	}
	return ans + accumulate(v[0])
}

//leetcode submit region end(Prohibit modification and deletion)
