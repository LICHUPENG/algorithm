package main

//你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。
//
// 注意：本题中，每个活字字模只能使用一次。
//
//
//
// 示例 1：
//
//
//输入："AAB"
//输出：8
//解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
//
//
// 示例 2：
//
//
//输入："AAABBC"
//输出：188
//
//
// 示例 3：
//
//
//输入："V"
//输出：1
//
//
//
// 提示：
//
//
// 1 <= tiles.length <= 7
// tiles 由大写英文字母组成
//
//
// Related Topics 哈希表 字符串 回溯 计数 👍 192 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func numTilePossibilities(tiles string) int {
	count := make(map[rune]int)
	for _, t := range tiles {
		count[t]++
	}
	return dfs(len(tiles), count) - 1
}

func dfs(i int, count map[rune]int) int {
	if i == 0 {
		return 1
	}
	res := 1
	for t := range count {
		if count[t] > 0 {
			count[t]--
			res += dfs(i-1, count)
			count[t]++
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
