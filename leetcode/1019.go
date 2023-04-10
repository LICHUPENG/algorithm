package main

//给定一个长度为 n 的链表 head
//
// 对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。
//
// 返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i 个节点没有下一个更大的节点
//，设置 answer[i] = 0 。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [2,1,5]
//输出：[5,5,0]
//
//
// 示例 2：
//
//
//
//
//输入：head = [2,7,4,3,5]
//输出：[7,0,5,5,0]
//
//
//
//
// 提示：
//
//
// 链表中节点数为 n
// 1 <= n <= 10⁴
// 1 <= Node.val <= 10⁹
//
//
// Related Topics 栈 数组 链表 单调栈 👍 264 👎 0

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	var ans []int
	var stack [][]int
	cur := head
	idx := -1
	for cur != nil {
		idx++
		ans = append(ans, 0)
		for len(stack) > 0 && stack[len(stack)-1][0] < cur.Val {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top[1]] = cur.Val
		}
		stack = append(stack, []int{cur.Val, idx})
		cur = cur.Next
	}
	return ans
}

func main() {
	println(nextLargerNodes(&ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 0,
					},
				},
			},
		},
	}))
}

//leetcode submit region end(Prohibit modification and deletion)
