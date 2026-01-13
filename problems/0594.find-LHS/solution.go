package find_LHS

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

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num] += 1
		} else {
			m[num] = 1
		}
	}

	res := 0
	for k, v := range m {
		if cnt, ok := m[k+1]; ok {
			res = max(res, cnt+v)
		}
	}

	return res
}
