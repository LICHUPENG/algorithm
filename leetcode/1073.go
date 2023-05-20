package main

//给出基数为 -2 的两个数 arr1 和 arr2，返回两数相加的结果。
//
// 数字以 数组形式 给出：数组由若干 0 和 1 组成，按最高有效位到最低有效位的顺序排列。例如，arr = [1,1,0,1] 表示数字 (-2)^3 +
// (-2)^2 + (-2)^0 = -3。数组形式 中的数字 arr 也同样不含前导零：即 arr == [0] 或 arr[0] == 1。
//
// 返回相同表示形式的 arr1 和 arr2 相加的结果。两数的表示形式为：不含前导零、由若干 0 和 1 组成的数组。
//
//
//
// 示例 1：
//
//
//输入：arr1 = [1,1,1,1,1], arr2 = [1,0,1]
//输出：[1,0,0,0,0]
//解释：arr1 表示 11，arr2 表示 5，输出表示 16 。
//
//
//
//
//
// 示例 2：
//
//
//输入：arr1 = [0], arr2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：arr1 = [0], arr2 = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
//
// 1 <= arr1.length, arr2.length <= 1000
// arr1[i] 和 arr2[i] 都是 0 或 1
// arr1 和 arr2 都没有前导0
//
//
// Related Topics 数组 数学 👍 74 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func addNegabinary(arr1 []int, arr2 []int) (ans []int) {
	i := len(arr1) - 1
	j := len(arr2) - 1
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		x := carry
		if i >= 0 {
			x += arr1[i]
		}
		if j >= 0 {
			x += arr2[j]
		}

		if x >= 2 {
			ans = append(ans, x-2)
			carry = -1
		} else if x >= 0 {
			ans = append(ans, x)
			carry = 0
		} else {
			ans = append(ans, 1)
			carry = 1
		}
		i--
		j--
	}
	for len(ans) > 1 && ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
