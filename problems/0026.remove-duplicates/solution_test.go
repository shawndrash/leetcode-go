package remove_duplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantLen  int
		wantNums []int // 期望去重后的前 k 个元素
	}{
		{
			name:     "示例1：有重复元素",
			nums:     []int{1, 1, 2},
			wantLen:  2,
			wantNums: []int{1, 2},
		},
		{
			name:     "示例2：多个重复元素",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantLen:  5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "无重复元素",
			nums:     []int{1, 2, 3, 4, 5},
			wantLen:  5,
			wantNums: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "所有元素相同",
			nums:     []int{1, 1, 1, 1, 1},
			wantLen:  1,
			wantNums: []int{1},
		},
		{
			name:     "单个元素",
			nums:     []int{1},
			wantLen:  1,
			wantNums: []int{1},
		},
		{
			name:     "两个相同元素",
			nums:     []int{1, 1},
			wantLen:  1,
			wantNums: []int{1},
		},
		{
			name:     "两个不同元素",
			nums:     []int{1, 2},
			wantLen:  2,
			wantNums: []int{1, 2},
		},
		{
			name:     "包含负数",
			nums:     []int{-3, -3, -2, -1, 0, 0, 1, 1},
			wantLen:  5,
			wantNums: []int{-3, -2, -1, 0, 1},
		},
		{
			name:     "长数组",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8},
			wantLen:  8,
			wantNums: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建输入数组的副本，因为函数会修改原数组
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			gotLen := removeDuplicates(nums)

			// 验证返回的长度
			if gotLen != tt.wantLen {
				t.Errorf("removeDuplicates() = %d, want %d", gotLen, tt.wantLen)
			}

			// 验证前 k 个元素是否正确去重
			gotNums := nums[:gotLen]
			if !reflect.DeepEqual(gotNums, tt.wantNums) {
				t.Errorf("removeDuplicates() nums[:%d] = %v, want %v", gotLen, gotNums, tt.wantNums)
			}

			// 验证去重后的元素确实不包含重复项
			for i := 1; i < gotLen; i++ {
				if gotNums[i] == gotNums[i-1] {
					t.Errorf("removeDuplicates() 去重后的数组包含重复项: nums[%d] = %d, nums[%d] = %d", i-1, gotNums[i-1], i, gotNums[i])
				}
			}
		})
	}
}

// TestRemoveDuplicatesEmpty 测试空数组
func TestRemoveDuplicatesEmpty(t *testing.T) {
	nums := []int{}
	gotLen := removeDuplicates(nums)
	if gotLen != 0 {
		t.Errorf("removeDuplicates([]) = %d, want 0", gotLen)
	}
	if len(nums) != 0 {
		t.Errorf("removeDuplicates([]) 修改了空数组")
	}
}

// BenchmarkSolution 性能测试
func BenchmarkSolution(b *testing.B) {
	// 创建一个较长的有序数组，包含重复元素
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i / 2 // 每两个元素相同，产生重复
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		removeDuplicates(testNums)
	}
}
