package main

import "strings"

//给定一个二进制字符串 s 和一个正整数 n，如果对于 [1, n] 范围内的每个整数，其二进制表示都是 s 的 子字符串 ，就返回 true，否则返回
//false 。
//
// 子字符串 是字符串中连续的字符序列。
//
//
//
// 示例 1：
//
//
//输入：s = "0110", n = 3
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "0110", n = 4
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s[i] 不是 '0' 就是 '1'
// 1 <= n <= 10⁹
//
//
// Related Topics 字符串 👍 90 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func help(s string, k int, mi int, ma int) bool {
	st := map[int]struct{}{}
	t := 0
	for r := 0; r < len(s); r++ {
		t = (t << 1) + int(s[r]-'0')
		if r >= k {
			t -= int(s[r-k]-'0') << k
		}
		if r >= k-1 && t >= mi && t <= ma {
			st[t] = struct{}{}
		}
	}
	return len(st) == ma-mi+1
}

func queryString(s string, n int) bool {
	if n == 1 {
		return strings.Contains(s, "1")
	}
	k := 30
	for (1 << k) >= n {
		k--
	}
	if len(s) < (1<<(k-1))+k-1 || len(s) < n-(1<<k)+k+1 {
		return false
	}
	return help(s, k, 1<<(k-1), (1<<k)-1) && help(s, k+1, 1<<k, n)
}

//leetcode submit region end(Prohibit modification and deletion)
