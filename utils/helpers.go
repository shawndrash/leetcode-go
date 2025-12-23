package utils

// Min 返回两个整数中的最小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max 返回两个整数中的最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Abs 返回整数的绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IntSliceEqual 比较两个整数切片是否相等
func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// IntSliceContains 检查切片是否包含指定元素
func IntSliceContains(slice []int, target int) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// ReverseIntSlice 反转整数切片
func ReverseIntSlice(slice []int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

// CopyIntSlice 复制整数切片
func CopyIntSlice(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	return result
}
