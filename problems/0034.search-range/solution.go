package search_range

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
func searchRange(nums []int, target int) []int {
	leftBorder, rightBorder := getLeftBorder(nums, target), getRightBorder(nums, target)
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}

	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}

	return []int{-1, -1}
}

func getLeftBorder(nums []int, target int) int {
	l, r, leftBorder := 0, len(nums)-1, -2
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
			leftBorder = r
		} else {
			l = mid + 1
		}
	}

	return leftBorder
}

func getRightBorder(nums []int, target int) int {
	l, r, rightBorder := 0, len(nums)-1, -2
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
			rightBorder = l
		}
	}

	return rightBorder
}
