package majority_element

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		// TODO: 添加测试参数
		want interface{}
	}{
		{
			name: "示例1",
			// TODO: 添加测试用例
			want: nil,
		},
		{
			name: "示例2",
			// TODO: 添加测试用例
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: 调用函数并验证结果
			// got := majorityUelement(...)
			// if got != tt.want {
			//     t.Errorf("majorityUelement() = %v, want %v", got, tt.want)
			// }
		})
	}
}

// BenchmarkSolution 性能测试
func BenchmarkSolution(b *testing.B) {
	// TODO: 添加性能测试
	for i := 0; i < b.N; i++ {
		// majorityUelement(...)
	}
}
