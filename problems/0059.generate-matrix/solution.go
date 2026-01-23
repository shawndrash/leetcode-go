package generate_matrix

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
func generateMatrix(n int) [][]int {
	matrix, m := make([][]int, n), 1
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	tR, tC, dR, dC := 0, 0, n-1, n-1
	fillEdge := func(matrix [][]int, tR, tC, dR, dC int) {
		curR, curC := tR, tC
		for curC <= dC {
			matrix[tR][curC] = m
			m++
			curC++
		}

		curR = tR + 1
		for curR <= dR {
			matrix[curR][dC] = m
			m++
			curR++
		}

		curC = dC - 1
		for curC >= tC {
			matrix[dR][curC] = m
			m++
			curC--
		}

		curR = dR - 1
		for curR > tR {
			matrix[curR][tC] = m
			m++
			curR--
		}
	}

	for tR <= dR {
		fillEdge(matrix, tR, tC, dR, dC)
		tR++
		tC++
		dR--
		dC--
	}

	return matrix
}
