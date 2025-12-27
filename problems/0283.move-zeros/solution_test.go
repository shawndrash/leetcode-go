package move_zeros

import (
	"fmt"
	"github.com/shawndrash/leetcode-go/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			nums: []int{0},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		moveZeroes(tt.nums)
		if !utils.IntSliceEqual(tt.nums, tt.want) {
			fmt.Printf("test fail, nums: %+v, want: %+v\n", tt.nums, tt.want)
		}
	}
}

// BenchmarkSolution 性能测试
func BenchmarkSolution(b *testing.B) {
	// TODO: 添加性能测试
	for i := 0; i < b.N; i++ {
		// moveUzeros(...)
	}
}
