package main

import "reflect"

//给你一个字符串数组 words ，每一个字符串长度都相同，令所有字符串的长度都为 n 。
//
// 每个字符串 words[i] 可以被转化为一个长度为 n - 1 的 差值整数数组 difference[i] ，其中对于 0 <= j <= n - 2
// 有 difference[i][j] = words[i][j+1] - words[i][j] 。注意两个字母的差值定义为它们在字母表中 位置 之差，也就是
//说 'a' 的位置是 0 ，'b' 的位置是 1 ，'z' 的位置是 25 。
//
//
// 比方说，字符串 "acb" 的差值整数数组是 [2 - 0, 1 - 2] = [2, -1] 。
//
//
// words 中所有字符串 除了一个字符串以外 ，其他字符串的差值整数数组都相同。你需要找到那个不同的字符串。
//
// 请你返回 words中 差值整数数组 不同的字符串。
//
//
//
// 示例 1：
//
//
//输入：words = ["adc","wzy","abc"]
//输出："abc"
//解释：
//- "adc" 的差值整数数组是 [3 - 0, 2 - 3] = [3, -1] 。
//- "wzy" 的差值整数数组是 [25 - 22, 24 - 25]= [3, -1] 。
//- "abc" 的差值整数数组是 [1 - 0, 2 - 1] = [1, 1] 。
//不同的数组是 [1, 1]，所以返回对应的字符串，"abc"。
//
//
// 示例 2：
//
//
//输入：words = ["aaa","bob","ccc","ddd"]
//输出："bob"
//解释：除了 "bob" 的差值整数数组是 [13, -13] 以外，其他字符串的差值整数数组都是 [0, 0] 。
//
//
//
//
// 提示：
//
//
// 3 <= words.length <= 100
// n == words[i].length
// 2 <= n <= 20
// words[i] 只含有小写英文字母。
//
//
// Related Topics 数组 哈希表 字符串 👍 40 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func get(word string) []int {
	diff := make([]int, len(word)-1)
	for i := 0; i+1 < len(word); i++ {
		diff[i] = int(word[i+1]) - int(word[i])
	}
	return diff
}

func oddString(words []string) string {
	diff0 := get(words[0])
	diff1 := get(words[1])
	if reflect.DeepEqual(diff0, diff1) {
		for i := 2; i < len(words); i++ {
			if !reflect.DeepEqual(diff0, get(words[i])) {
				return words[i]
			}
		}
	}
	if reflect.DeepEqual(diff0, get(words[2])) {
		return words[1]
	}
	return words[0]
}

//leetcode submit region end(Prohibit modification and deletion)
