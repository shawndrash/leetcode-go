package remove_element

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
func removeElement(nums []int, val int) int {
	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}
	l, r, cnt := 0, len(nums)-1, 0
	for l < r {
		for r >= 0 && nums[r] == val {
			r--
			cnt++
		}

		for l < len(nums) && nums[l] != val && l < r {
			l++
		}

		if r >= 0 && l < len(nums) {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	return len(nums) - cnt
}

func removeElement2(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if val != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
