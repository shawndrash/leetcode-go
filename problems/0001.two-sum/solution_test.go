package two_sum

import (
	"fmt"
	"github.com/shawndrash/leetcode-go/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !utils.IntSliceEqual(got, tt.want) {
			fmt.Printf("test fail, nums: %v, target:%v, got: %v, want: %v\n", tt.nums, tt.target, got, tt.want)
		}
	}
}
