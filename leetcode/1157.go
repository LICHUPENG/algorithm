package main

import (
	"math/rand"
	"sort"
	"time"
)

// 设计一个数据结构，有效地找到给定子数组的 多数元素 。
//
// 子数组的 多数元素 是在子数组中出现 threshold 次数或次数以上的元素。
//
// 实现 MajorityChecker 类:
//
// MajorityChecker(int[] arr) 会用给定的数组 arr 对 MajorityChecker 初始化。
// int query(int left, int right, int threshold) 返回子数组中的元素 arr[left...right] 至少出
// 现 threshold 次数，如果不存在这样的元素则返回 -1。
//
// 示例 1：
//
// 输入:
// ["MajorityChecker", "query", "query", "query"]
// [[[1, 1, 2, 2, 1, 1]], [0, 5, 4], [0, 3, 3], [2, 3, 2]]
// 输出：
// [null, 1, -1, 2]
//
// 解释：
// MajorityChecker majorityChecker = new MajorityChecker([1,1,2,2,1,1]);
// majorityChecker.query(0,5,4); // 返回 1
// majorityChecker.query(0,3,3); // 返回 -1
// majorityChecker.query(2,3,2); // 返回 2
//
// 提示：
//
// 1 <= arr.length <= 2 * 10⁴
// 1 <= arr[i] <= 2 * 10⁴
// 0 <= left <= right < arr.length
// threshold <= right - left + 1
// 2 * threshold > right - left + 1
// 调用 query 的次数最多为 10⁴
//
// Related Topics 设计 树状数组 线段树 数组 二分查找 👍 138 👎 0
func main() {
	rand.Seed(time.Now().UnixNano())
	arr := []int{1, 1, 2, 2, 1, 1}
	majorityChecker := Constructor(arr)
	print(majorityChecker.Query(0, 5, 4))
	print(majorityChecker.Query(0, 3, 3))
	print(majorityChecker.Query(2, 3, 2))
}

// leetcode submit region begin(Prohibit modification and deletion)
type MajorityChecker struct {
	arr []int
	loc map[int][]int
}

func Constructor(arr []int) MajorityChecker {
	rand.Seed(time.Now().UnixNano())
	loc := map[int][]int{}
	for i, x := range arr {
		loc[x] = append(loc[x], i)
	}
	return MajorityChecker{arr, loc}
}

func (mc *MajorityChecker) Query(left, right, threshold int) int {
	length := right - left + 1
	for i := 0; i < 20; i++ {
		x := mc.arr[rand.Intn(right-left+1)+left]
		pos := mc.loc[x]
		occ := sort.SearchInts(pos, right+1) - sort.SearchInts(pos, left)
		if occ >= threshold {
			return x
		}
		if occ*2 >= length {
			break
		}
	}
	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
//leetcode submit region end(Prohibit modification and deletion)
