package main

//给你一个整数 n ，以二进制字符串的形式返回该整数的 负二进制（base -2）表示。
//
// 注意，除非字符串就是 "0"，否则返回的字符串中不能含有前导零。
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出："110"
//解释：(-2)² + (-2)¹ = 2
//
//
// 示例 2：
//
//
//输入：n = 3
//输出："111"
//解释：(-2)² + (-2)¹ + (-2)⁰ = 3
//
//
// 示例 3：
//
//
//输入：n = 4
//输出："100"
//解释：(-2)² = 4
//
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁹
//
//
// Related Topics 数学 👍 153 👎 0

func main() {
	println(baseNeg2(3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func baseNeg2(n int) string {
	val := 0x55555555 ^ (0x55555555 - n)
	if val == 0 {
		return "0"
	}
	res := []byte{}
	for val > 0 {
		res = append(res, '0'+byte(val&1))
		val >>= 1
	}
	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return string(res)

}

//leetcode submit region end(Prohibit modification and deletion)
