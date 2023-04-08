package main

import "math/bits"

// 作为项目经理，你规划了一份需求的技能清单 req_skills，并打算从备选人员名单 people 中选出些人组成一个「必要团队」（ 编号为 i 的备选人员 people[i] 含有一份该备选人员掌握的技能列表）。
//
// 所谓「必要团队」，就是在这个团队中，对于所需求的技能列表 req_skills 中列出的每项技能，团队中至少有一名成员已经掌握。可以用每个人的编号来表示团队中的成员：
//
// 例如，团队 team = [0, 1, 3] 表示掌握技能分别为 people[0]，people[1]，和 people[3] 的备选人员。
// 请你返回 任一 规模最小的必要团队，团队成员用人员编号表示。你可以按 任意顺序 返回答案，题目数据保证答案存在。
func smallestSufficientTeam(reqSkills []string, people [][]string) (ans []int) {
	m := len(reqSkills)
	sid := make(map[string]int, m)
	for i, s := range reqSkills {
		sid[s] = i // 字符串映射到下标
	}

	n := len(people)
	mask := make([]int, n)
	for i, skills := range people {
		for _, s := range skills { // 把 skills 压缩成一个二进制数 mask[i]
			mask[i] |= 1 << sid[s]
		}
	}

	u, all := 1<<m, 1<<n-1
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, u)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示还没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j == 0 { // 背包已装满
			return 0
		}
		if i < 0 { // 没法装满背包，返回全集，这样下面比较集合大小会取更小的
			return all
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		r1 := dfs(i-1, j)                 // 不选 mask[i]
		r2 := dfs(i-1, j&^mask[i]) | 1<<i // 选 mask[i]
		if bits.OnesCount(uint(r1)) < bits.OnesCount(uint(r2)) {
			*p = r1
		} else {
			*p = r2
		}
		return *p
	}
	res := dfs(n-1, u-1)

	for i := 0; i < n; i++ {
		if res>>i&1 > 0 {
			ans = append(ans, i) // 所有在 res 中的下标
		}
	}
	return
}
