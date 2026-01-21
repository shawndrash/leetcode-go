package total_fruit

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
func totalFruit(fruits []int) int {
	m := make(map[int]int, 2)
	i, res := 0, 1
	m[fruits[i]] = i
	for j := 1; j < len(fruits); j++ {
		m[fruits[j]] = j
		if len(m) <= 2 {
			res = max(res, j-i+1)
		} else {
			minIndex := len(fruits)
			for _, v := range m {
				minIndex = min(minIndex, v)
			}

			delete(m, fruits[minIndex])
			i = minIndex + 1
		}
	}

	return res
}
