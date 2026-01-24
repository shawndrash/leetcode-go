package four_sum_count

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

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			m[a+b]++
		}
	}

	res := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			if count, ok := m[0-c-d]; ok {
				res += count
			}
		}
	}

	return res
}
