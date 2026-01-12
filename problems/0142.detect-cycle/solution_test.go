package detect_cycle

import (
	"testing"
)

// buildList 从切片构建链表
func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	return head
}

// createCycle 创建带环的链表
// values: 链表的值
// pos: 环的入口位置（-1 表示无环，0 表示第一个节点，以此类推）
func createCycle(values []int, pos int) (*ListNode, *ListNode) {
	if len(values) == 0 {
		return nil, nil
	}
	head := buildList(values)
	if pos == -1 {
		return head, nil
	}

	// 找到环的入口节点
	var entry *ListNode
	current := head
	for i := 0; current != nil; i++ {
		if i == pos {
			entry = current
		}
		if current.Next == nil {
			// 将最后一个节点指向入口节点
			current.Next = entry
			break
		}
		current = current.Next
	}
	return head, entry
}

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name    string
		values  []int
		pos     int // 环的入口位置，-1 表示无环
		wantVal int // 期望返回节点的值（nil 时使用 -1 表示）
		wantNil bool
	}{
		{
			name:    "无环链表",
			values:  []int{1, 2, 3, 4},
			pos:     -1,
			wantNil: true,
		},
		{
			name:    "空链表",
			values:  []int{},
			pos:     -1,
			wantNil: true,
		},
		{
			name:    "单个节点无环",
			values:  []int{1},
			pos:     -1,
			wantNil: true,
		},
		{
			name:    "环在第一个节点",
			values:  []int{3, 2, 0, -4},
			pos:     0,
			wantVal: 3,
			wantNil: false,
		},
		{
			name:    "环在第二个节点",
			values:  []int{3, 2, 0, -4},
			pos:     1,
			wantVal: 2,
			wantNil: false,
		},
		{
			name:    "环在最后一个节点（指向自己）",
			values:  []int{1, 2},
			pos:     1,
			wantVal: 2,
			wantNil: false,
		},
		{
			name:    "两个节点形成环",
			values:  []int{1, 2},
			pos:     0,
			wantVal: 1,
			wantNil: false,
		},
		{
			name:    "环在中间节点",
			values:  []int{1, 2, 3, 4, 5},
			pos:     2,
			wantVal: 3,
			wantNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head, expectedEntry := createCycle(tt.values, tt.pos)
			got := detectCycle(head)

			if tt.wantNil {
				if got != nil {
					t.Errorf("detectCycle() = %v, want nil", got)
				}
			} else {
				if got == nil {
					t.Errorf("detectCycle() = nil, want node with value %d", tt.wantVal)
					return
				}
				if got != expectedEntry {
					t.Errorf("detectCycle() = node with value %d, want node with value %d", got.Val, tt.wantVal)
				}
				if got.Val != tt.wantVal {
					t.Errorf("detectCycle() = node with value %d, want node with value %d", got.Val, tt.wantVal)
				}
			}
		})
	}
}

// BenchmarkSolution 性能测试
func BenchmarkSolution(b *testing.B) {
	// 创建一个较长的带环链表
	values := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		values[i] = i
	}
	head, _ := createCycle(values, 500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detectCycle(head)
	}
}
