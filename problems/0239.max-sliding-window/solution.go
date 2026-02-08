package max_sliding_window

// 如果是链表题目，取消下面的注释：
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 如果是树题目，取消下面的注释：
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// ============= 以下是提交到 LeetCode 的代码 =============
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{make([]int, 0)}
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.queue[0] {
		m.queue = m.queue[1:]
	}
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}

	m.queue = append(m.queue, val)
}

func maxSlidingWindow(nums []int, k int) []int {
	queue, res := NewMyQueue(), make([]int, 0, len(nums)-k+1)
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}

	res = append(res, queue.Front())
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		res = append(res, queue.Front())
	}

	return res
}
