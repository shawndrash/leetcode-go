package can_complete_circuit

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
func canCompleteCircuit(gas []int, cost []int) int {
	restSum, curSum, res := 0, 0, -1
	for i := 0; i < len(gas); i++ {
		restSum += gas[i] - cost[i]
		curSum += gas[i] - cost[i]
		if curSum < 0 {
			res = i + 1
			curSum = 0
		}
	}

	if restSum < 0 {
		return -1
	}

	return res
}
