package next_greater_element

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
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res, stack, m := make([]int, len(nums1)), make([]int, 0, len(nums2)), make(map[int]int)
	for i := range res {
		res[i] = -1
	}

	for i, num := range nums1 {
		m[num] = i
	}

	stack = append(stack, 0)
	for i := 1; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			if index, ok := m[nums2[stack[len(stack)-1]]]; ok {
				res[index] = nums2[i]
			}

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
