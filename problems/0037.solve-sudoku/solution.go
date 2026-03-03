package solve_sudoku

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
func solveSudoku(board [][]byte) {
	var backtrack func() bool
	backtrack = func() bool {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] == '.' {
					continue
				}
				for k := '1'; k <= '9'; k++ {
					num := byte(k)
					if isValid(board, i, j, num) {
						board[i][j] = num
						if backtrack() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}

		return true
	}

	backtrack()
}

func isValid(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] == num {
			return false
		}
	}

	for j := 0; j < len(board[0]); j++ {
		if board[row][j] == num {
			return false
		}
	}

	startRow, startCol := row/3*3, col/3*3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}
