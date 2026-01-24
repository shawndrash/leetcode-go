package intersection

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
func intersection(nums1 []int, nums2 []int) []int {
	res, m := make([]int, 0), make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = struct{}{}
	}

	resSet := make(map[int]struct{})
	for _, num2 := range nums2 {
		if _, ok := m[num2]; ok {
			resSet[num2] = struct{}{}
		}
	}

	for k, _ := range resSet {
		res = append(res, k)
	}

	return res
}
