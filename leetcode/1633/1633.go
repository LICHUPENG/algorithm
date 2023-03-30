package main

import "sort"

func main() {
	println(maxWidthOfVerticalArea([][]int{
		{8, 7},
		{8, 9},
		{9, 7},
		{7, 4}}))
}

func maxWidthOfVerticalArea(points [][]int) (ans int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	for i := 1; i < len(points); i++ {
		ans = max(ans, points[i][0]-points[i-1][0])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
