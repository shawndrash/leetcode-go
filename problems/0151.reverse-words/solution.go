package reverse_words

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
func reverse(b []byte) {
	l, r := 0, len(b)-1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}

func reverseWords(s string) string {
	slow, fast, b := 0, 0, []byte(s)
	for fast < len(b) && b[fast] == ' ' {
		fast++
	}

	for ; fast < len(b); fast++ {
		if fast-1 > 0 && b[fast-1] == b[fast] && b[fast] == ' ' {
			continue
		}

		b[slow] = b[fast]
		slow++
	}

	b = b[:slow]
	if slow-1 > 0 && b[slow-1] == ' ' {
		b = b[:slow-1]
	}

	reverse(b)

	slow, fast = 0, 0
	for fast < len(b) {
		if b[fast] == ' ' {
			reverse(b[slow:fast])
			slow = fast + 1
		}

		fast++
	}

	reverse(b[slow:])
	return string(b)
}
