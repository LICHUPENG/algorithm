package main

//我们把无限数量 ∞ 的栈排成一行，按从左到右的次序从 0 开始编号。每个栈的的最大容量 capacity 都相同。
//
// 实现一个叫「餐盘」的类 DinnerPlates：
//
//
// DinnerPlates(int capacity) - 给出栈的最大容量 capacity。
// void push(int val) - 将给出的正整数 val 推入 从左往右第一个 没有满的栈。
// int pop() - 返回 从右往左第一个 非空栈顶部的值，并将其从栈中删除；如果所有的栈都是空的，请返回 -1。
// int popAtStack(int index) - 返回编号 index 的栈顶部的值，并将其从栈中删除；如果编号 index 的栈是空的，请返回 -
//1。
//
//
//
//
// 示例：
//
// 输入：
//["DinnerPlates","push","push","push","push","push","popAtStack","push","push",
//"popAtStack","popAtStack","pop","pop","pop","pop","pop"]
//[[2],[1],[2],[3],[4],[5],[0],[20],[21],[0],[2],[],[],[],[],[]]
//输出：
//[null,null,null,null,null,null,2,null,null,20,21,5,4,3,1,-1]
//
//解释：
//DinnerPlates D = DinnerPlates(2);  // 初始化，栈最大容量 capacity = 2
//D.push(1);
//D.push(2);
//D.push(3);
//D.push(4);
//D.push(5);         // 栈的现状为：    2  4
//                                    1  3  5
//                                    ﹈ ﹈ ﹈
//D.popAtStack(0);   // 返回 2。栈的现状为：      4
//                                          1  3  5
//                                          ﹈ ﹈ ﹈
//D.push(20);        // 栈的现状为：  20  4
//                                   1  3  5
//                                   ﹈ ﹈ ﹈
//D.push(21);        // 栈的现状为：  20  4 21
//                                   1  3  5
//                                   ﹈ ﹈ ﹈
//D.popAtStack(0);   // 返回 20。栈的现状为：       4 21
//                                            1  3  5
//                                            ﹈ ﹈ ﹈
//D.popAtStack(2);   // 返回 21。栈的现状为：       4
//                                            1  3  5
//                                            ﹈ ﹈ ﹈
//D.pop()            // 返回 5。栈的现状为：        4
//                                            1  3
//                                            ﹈ ﹈
//D.pop()            // 返回 4。栈的现状为：    1  3
//                                           ﹈ ﹈
//D.pop()            // 返回 3。栈的现状为：    1
//                                           ﹈
//D.pop()            // 返回 1。现在没有栈。
//D.pop()            // 返回 -1。仍然没有栈。
//
//
//
//
// 提示：
//
//
// 1 <= capacity <= 20000
// 1 <= val <= 20000
// 0 <= index <= 100000
// 最多会对 push，pop，和 popAtStack 进行 200000 次调用。
//
//
// Related Topics 栈 设计 哈希表 堆（优先队列） 👍 86 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
type DinnerPlates struct {
	capacity  int
	stack     []int
	top       []int
	poppedPos []int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{capacity, []int{}, []int{}, []int{}}
}

func (this *DinnerPlates) Push(val int) {
	if len(this.poppedPos) == 0 {
		pos := len(this.stack)
		this.stack = append(this.stack, val)
		if pos%this.capacity == 0 {
			this.top = append(this.top, 0)
		} else {
			stackPos := len(this.top) - 1
			stackTop := this.top[stackPos]
			this.top[stackPos] = stackTop + 1
		}
	} else {
		pos := this.poppedPos[0]
		this.poppedPos = this.poppedPos[1:]
		this.stack[pos] = val
		index := pos / this.capacity
		stackTop := this.top[index]
		this.top[index] = stackTop + 1
	}
}

func (this *DinnerPlates) Pop() int {
	for len(this.stack) > 0 && len(this.poppedPos) > 0 && this.poppedPos[len(this.poppedPos)-1] == len(this.stack)-1 {
		this.stack = this.stack[:len(this.stack)-1]
		pos := this.poppedPos[len(this.poppedPos)-1]
		this.poppedPos = this.poppedPos[:len(this.poppedPos)-1]
		if pos%this.capacity == 0 {
			this.top = this.top[:len(this.top)-1]
		}
	}
	if len(this.stack) == 0 {
		return -1
	} else {
		pos := len(this.stack) - 1
		val := this.stack[pos]
		this.stack = this.stack[:pos]
		if pos%this.capacity == 0 && len(this.top) > 0 {
			this.top = this.top[:len(this.top)-1]
		} else if len(this.top) > 0 {
			this.top[len(this.top)-1] -= 1
		}
		return val
	}
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index >= len(this.top) {
		return -1
	}
	stackTop := this.top[index]
	if stackTop < 0 {
		return -1
	}
	this.top[index] = stackTop - 1
	pos := index*this.capacity + stackTop
	i := sort.SearchInts(this.poppedPos, pos)
	this.poppedPos = append(this.poppedPos, 0)
	copy(this.poppedPos[i+1:], this.poppedPos[i:])
	this.poppedPos[i] = pos
	return this.stack[pos]
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */
//leetcode submit region end(Prohibit modification and deletion)
