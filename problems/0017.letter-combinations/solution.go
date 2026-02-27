package letter_combinations

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
func letterCombinations(digits string) []string {
	mapping := map[uint8]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	res, path := make([]string, 0), ""
	var backtrack func(digitIndex int)
	backtrack = func(digitIndex int) {
		if len(path) == len(digits) {
			res = append(res, path)
			return
		}

		letters := mapping[digits[digitIndex]]
		for i := 0; i < len(letters); i++ {
			path += string(letters[i])
			backtrack(digitIndex + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
