package remove_element

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantLen  int
		wantNums []int // 期望移除后的前 k 个元素（顺序可能不同，但不应包含 val）
	}{
		{
			name:     "示例1：移除值为3的元素",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			wantLen:  2,
			wantNums: []int{2, 2},
		},
		{
			name:     "示例2：移除值为2的元素",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			wantLen:  5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
		{
			name:     "无目标元素",
			nums:     []int{1, 2, 3, 4, 5},
			val:      6,
			wantLen:  5,
			wantNums: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "所有元素都是目标值",
			nums:     []int{3, 3, 3, 3},
			val:      3,
			wantLen:  0,
			wantNums: []int{},
		},
		{
			name:     "单个元素且为目标值",
			nums:     []int{1},
			val:      1,
			wantLen:  0,
			wantNums: []int{},
		},
		{
			name:     "单个元素但不是目标值",
			nums:     []int{1},
			val:      2,
			wantLen:  1,
			wantNums: []int{1},
		},
		{
			name:     "空数组",
			nums:     []int{},
			val:      1,
			wantLen:  0,
			wantNums: []int{},
		},
		{
			name:     "移除值为0的元素",
			nums:     []int{0, 1, 0, 3, 12},
			val:      0,
			wantLen:  3,
			wantNums: []int{1, 3, 12},
		},
		{
			name:     "包含负数",
			nums:     []int{-1, 0, 1, -1, 2, -1},
			val:      -1,
			wantLen:  3,
			wantNums: []int{0, 1, 2},
		},
		{
			name:     "目标值在开头",
			nums:     []int{2, 2, 2, 1, 3, 4},
			val:      2,
			wantLen:  3,
			wantNums: []int{1, 3, 4},
		},
		{
			name:     "目标值在末尾",
			nums:     []int{1, 3, 4, 2, 2, 2},
			val:      2,
			wantLen:  3,
			wantNums: []int{1, 3, 4},
		},
		{
			name:     "目标值在中间",
			nums:     []int{1, 2, 2, 3, 4},
			val:      2,
			wantLen:  3,
			wantNums: []int{1, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建输入数组的副本，因为函数会修改原数组
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			gotLen := removeElement(nums, tt.val)

			// 验证返回的长度
			if gotLen != tt.wantLen {
				t.Errorf("removeElement() = %d, want %d", gotLen, tt.wantLen)
			}

			// 验证前 k 个元素不包含目标值
			gotNums := nums[:gotLen]
			for i, num := range gotNums {
				if num == tt.val {
					t.Errorf("removeElement() 前 %d 个元素中包含目标值 %d: nums[%d] = %d", gotLen, tt.val, i, num)
				}
			}

			// 验证前 k 个元素的数量正确（不包含目标值的元素数量）
			expectedCount := 0
			for _, num := range tt.nums {
				if num != tt.val {
					expectedCount++
				}
			}
			if gotLen != expectedCount {
				t.Errorf("removeElement() 返回长度 %d 不正确，期望移除后应有 %d 个元素", gotLen, expectedCount)
			}

			// 验证前 k 个元素确实不包含目标值（通过统计）
			valCount := 0
			for _, num := range gotNums {
				if num == tt.val {
					valCount++
				}
			}
			if valCount > 0 {
				t.Errorf("removeElement() 前 %d 个元素中包含 %d 个目标值 %d", gotLen, valCount, tt.val)
			}
		})
	}
}

// TestRemoveElementOrder 测试移除后元素的顺序（虽然题目不要求顺序，但可以验证）
func TestRemoveElementOrder(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	originalNums := make([]int, len(nums))
	copy(originalNums, nums)

	gotLen := removeElement(nums, val)

	// 验证移除后的元素确实来自原数组
	result := nums[:gotLen]
	originalNonVal := []int{}
	for _, num := range originalNums {
		if num != val {
			originalNonVal = append(originalNonVal, num)
		}
	}

	// 验证结果中的元素都在原数组的非目标值元素中
	resultMap := make(map[int]int)
	for _, num := range result {
		resultMap[num]++
	}

	originalMap := make(map[int]int)
	for _, num := range originalNonVal {
		originalMap[num]++
	}

	// 验证每个元素的数量匹配
	for num, count := range resultMap {
		if originalMap[num] != count {
			t.Errorf("removeElement() 结果中元素 %d 的数量 %d 与原数组不匹配，期望 %d", num, count, originalMap[num])
		}
	}
}

// BenchmarkSolution 性能测试
func BenchmarkSolution(b *testing.B) {
	// 创建一个较长的数组，包含要移除的元素
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			nums[i] = 2 // 要移除的值
		} else {
			nums[i] = i
		}
	}
	val := 2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		removeElement(testNums, val)
	}
}
