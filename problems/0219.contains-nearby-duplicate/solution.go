package contains_nearby_duplicate

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

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]struct{})
	for i, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}

		m[num] = struct{}{}
		if len(m) > k {
			delete(m, i-k)
		}
	}

	return false
}
