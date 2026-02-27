package restore_ip_addresses

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
func restoreIpAddresses(s string) []string {
	res, dotNum := make([]string, 0), 0
	var backtrack func(start int)
	backtrack = func(start int) {
		if dotNum == 3 {
			if isValid(s, start, len(s)-1) {
				res = append(res, s)
			}
			return
		}

		for i := start; i < len(s); i++ {
			if isValid(s, start, i) {
				s = s[:i+1] + "." + s[i+1:]
				dotNum++
				backtrack(i + 2)
				dotNum--
				s = s[:i+1] + s[i+2:]
			}
		}
	}

	backtrack(0)
	return res
}

func isValid(s string, start, end int) bool {
	if end-start > 2 || start >= len(s) {
		return false
	}
	if s[start] == '0' && start != end {
		return false
	}
	num := 0
	for i := start; i <= end; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		num = num*10 + int(s[i]-'0')
		if num > 255 {
			return false
		}
	}

	return true
}
