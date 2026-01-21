package move_zeros

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

// TODO: 实现你的解法
// 提交时，直接复制下面的函数（及上面需要的结构体定义）到 LeetCode 即可
func moveZeroes(nums []int) {
	l := 0
	for l < len(nums) && nums[l] != 0 {
		l++
	}

	if l >= len(nums) {
		return
	}

	for r := l + 1; r < len(nums); r++ {
		if nums[r] == 0 {
			continue
		}

		nums[l], nums[r] = nums[r], nums[l]
		l++
	}
}

func moveZeroes2(nums []int) {
	slow := 0
	for slow < len(nums) && nums[slow] != 0 {
		slow++
	}

	for fast := slow + 1; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}

	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}
