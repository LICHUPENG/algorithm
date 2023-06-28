package main

//给你一个 n 个节点的 无向带权连通 图，节点编号为 0 到 n - 1 ，再给你一个整数数组 edges ，其中 edges[i] = [ai, bi,
//wi] 表示节点 ai 和 bi 之间有一条边权为 wi 的边。
//
// 部分边的边权为 -1（wi = -1），其他边的边权都为 正 数（wi > 0）。
//
// 你需要将所有边权为 -1 的边都修改为范围 [1, 2 * 10⁹] 中的 正整数 ，使得从节点 source 到节点 destination 的 最短距
//离 为整数 target 。如果有 多种 修改方案可以使 source 和 destination 之间的最短距离等于 target ，你可以返回任意一种方案。
//
//
// 如果存在使 source 到 destination 最短距离为 target 的方案，请你按任意顺序返回包含所有边的数组（包括未修改边权的边）。如果不存
//在这样的方案，请你返回一个 空数组 。
//
// 注意：你不能修改一开始边权为正数的边。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 5, edges = [[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]], source = 0,
//destination = 1, target = 5
//输出：[[4,1,1],[2,0,1],[0,3,3],[4,3,1]]
//解释：上图展示了一个满足题意的修改方案，从 0 到 1 的最短距离为 5 。
//
//
// 示例 2：
//
//
//
//
//输入：n = 3, edges = [[0,1,-1],[0,2,5]], source = 0, destination = 2, target = 6
//输出：[]
//解释：上图是一开始的图。没有办法通过修改边权为 -1 的边，使得 0 到 2 的最短距离等于 6 ，所以返回一个空数组。
//
//
// 示例 3：
//
//
//
//
//输入：n = 4, edges = [[1,0,4],[1,2,3],[2,3,5],[0,3,-1]], source = 0, destination
//= 2, target = 6
//输出：[[1,0,4],[1,2,3],[2,3,5],[0,3,1]]
//解释：上图展示了一个满足题意的修改方案，从 0 到 2 的最短距离为 6 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// 1 <= edges.length <= n * (n - 1) / 2
// edges[i].length == 3
// 0 <= ai, bi < n
// wi = -1 或者 1 <= wi <= 107
// ai != bi
// 0 <= source, destination < n
// source != destination
// 1 <= target <= 10⁹
// 输入的图是连通图，且没有自环和重边。
//
//
// Related Topics 图 最短路 堆（优先队列） 👍 55 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	k := 0
	for _, e := range edges {
		if e[2] == -1 {
			k++
		}
	}
	if dijkstra(source, destination, construct(n, edges, 0, target)) > int64(target) {
		return nil
	}
	if dijkstra(source, destination, construct(n, edges, int64(k)*int64(target-1), target)) < int64(target) {
		return nil
	}

	left, right, ans := int64(0), int64(k)*int64(target-1), int64(0)
	for left <= right {
		mid := int64(left+right) / 2
		if dijkstra(source, destination, construct(n, edges, mid, target)) >= int64(target) {
			ans, right = mid, mid-1
		} else {
			left = mid + 1
		}
	}
	for _, e := range edges {
		if e[2] == -1 {
			if ans >= int64(target-1) {
				e[2] = target
				ans -= int64(target - 1)
			} else {
				e[2] = int(1 + ans)
				ans = 0
			}
		}
	}
	return edges
}

func dijkstra(source, destination int, adjMatrix [][]int) int64 {
	// 朴素的 dijistra 算法
	// adjMatrix 是一个邻接矩阵
	n := len(adjMatrix)
	dist, used := make([]int64, n), make([]bool, n)
	for i := 0; i < n; i++ {
		dist[i] = 0x3f3f3f3f3f
	}
	dist[source] = 0
	for round := 0; round < n-1; round++ {
		u := -1
		for i := 0; i < n; i++ {
			if !used[i] && (u == -1 || dist[i] < dist[u]) {
				u = i
			}
		}
		used[u] = true
		for v := 0; v < n; v++ {
			if !used[v] && adjMatrix[u][v] != -1 && dist[v] > dist[u]+int64(adjMatrix[u][v]) {
				dist[v] = dist[u] + int64(adjMatrix[u][v])
			}
		}
	}
	return dist[destination]
}

func construct(n int, edges [][]int, idx int64, target int) [][]int {
	// 需要构造出第 idx 种不同的边权情况，返回一个邻接矩阵
	adjMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		adjMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			adjMatrix[i][j] = -1
		}
	}
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if w != -1 {
			adjMatrix[u][v], adjMatrix[v][u] = w, w
		} else {
			if idx >= int64(target-1) {
				adjMatrix[u][v], adjMatrix[v][u] = target, target
				idx -= int64(target - 1)
			} else {
				adjMatrix[u][v], adjMatrix[v][u] = int(1+idx), int(1+idx)
				idx = 0
			}
		}
	}
	return adjMatrix
}

//leetcode submit region end(Prohibit modification and deletion)
