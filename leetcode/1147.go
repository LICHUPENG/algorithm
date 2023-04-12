package main

//你会得到一个字符串 text 。你应该把它分成 k 个子字符串 (subtext1, subtext2，…， subtextk) ，要求满足:
//
//
// subtexti 是 非空 字符串
// 所有子字符串的连接等于 text ( 即subtext1 + subtext2 + ... + subtextk == text )
// 对于所有 i 的有效值( 即 1 <= i <= k ) ，subtexti == subtextk - i + 1 均成立
//
//
// 返回k可能最大值。
//
//
//
// 示例 1：
//
//
//输入：text = "ghiabcdefhelloadamhelloabcdefghi"
//输出：7
//解释：我们可以把字符串拆分成 "(ghi)(abcdef)(hello)(adam)(hello)(abcdef)(ghi)"。
//
//
// 示例 2：
//
//
//输入：text = "merchant"
//输出：1
//解释：我们可以把字符串拆分成 "(merchant)"。
//
//
// 示例 3：
//
//
//输入：text = "antaprezatepzapreanta"
//输出：11
//解释：我们可以把字符串拆分成 "(a)(nt)(a)(pre)(za)(tpe)(za)(pre)(a)(nt)(a)"。
//
//
//
//
// 提示：
//
//
// 1 <= text.length <= 1000
// text 仅由小写英文字符组成
//
//
// Related Topics 贪心 双指针 字符串 动态规划 哈希函数 滚动哈希 👍 129 👎 0

func main() {
	println(longestDecomposition("ghiabcdefhelloadamhelloabcdefghi"))
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestDecomposition(s string) int {
	if s == "" {
		return 0
	}
	for i, n := 1, len(s); i <= n/2; i++ { // 枚举前后缀长度
		if s[:i] == s[n-i:] { // 立刻分割
			return 2 + longestDecomposition(s[i:n-i])
		}
	}
	return 1 // 无法分割
}

//leetcode submit region end(Prohibit modification and deletion)
