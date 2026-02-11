package reverse_words

import (
	"reflect"
	"testing"
)

// TestReverse 测试 reverse 函数
func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "正常字节切片",
			input:    []byte("hello"),
			expected: []byte("olleh"),
		},
		{
			name:     "两个字符",
			input:    []byte("ab"),
			expected: []byte("ba"),
		},
		{
			name:     "单个字符",
			input:    []byte("a"),
			expected: []byte("a"),
		},
		{
			name:     "空切片",
			input:    []byte(""),
			expected: []byte(""),
		},
		{
			name:     "包含空格",
			input:    []byte("a b c"),
			expected: []byte("c b a"),
		},
		{
			name:     "数字",
			input:    []byte("12345"),
			expected: []byte("54321"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建输入副本，因为函数会修改原切片
			input := make([]byte, len(tt.input))
			copy(input, tt.input)

			reverse(input)

			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("reverse() = %v, want %v", string(input), string(tt.expected))
			}
		})
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "示例1：正常单词",
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			name:     "示例2：前后有空格",
			input:    "  hello world  ",
			expected: "world hello",
		},
		{
			name:     "示例3：多个空格",
			input:    "a good   example",
			expected: "example good a",
		},
		{
			name:     "单个单词",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "单个单词前后有空格",
			input:    "  hello  ",
			expected: "hello",
		},
		{
			name:     "两个单词",
			input:    "hello world",
			expected: "world hello",
		},
		{
			name:     "三个单词",
			input:    "one two three",
			expected: "three two one",
		},
		{
			name:     "只有空格",
			input:    "   ",
			expected: "",
		},
		{
			name:     "空字符串",
			input:    "",
			expected: "",
		},
		{
			name:     "单词间多个空格",
			input:    "a    b     c",
			expected: "c b a",
		},
		{
			name:     "开头多个空格",
			input:    "   hello world",
			expected: "world hello",
		},
		{
			name:     "结尾多个空格",
			input:    "hello world   ",
			expected: "world hello",
		},
		{
			name:     "前后都有多个空格",
			input:    "   hello   world   ",
			expected: "world hello",
		},
		{
			name:     "包含标点符号",
			input:    "hello, world!",
			expected: "world! hello,",
		},
		{
			name:     "数字和字母",
			input:    "test 123 abc",
			expected: "abc 123 test",
		},
		{
			name:     "长句子",
			input:    "the quick brown fox jumps over the lazy dog",
			expected: "dog lazy the over jumps fox brown quick the",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.input)
			if got != tt.expected {
				t.Errorf("reverseWords(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// TestReverseWordsEdgeCases 测试边界情况
func TestReverseWordsEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "单个字符",
			input:    "a",
			expected: "a",
		},
		{
			name:     "单个字符前后空格",
			input:    "  a  ",
			expected: "a",
		},
		{
			name:     "相同单词",
			input:    "test test test",
			expected: "test test test",
		},
		{
			name:     "单词长度不同",
			input:    "a bb ccc dddd",
			expected: "dddd ccc bb a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.input)
			if got != tt.expected {
				t.Errorf("reverseWords(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// TestReverseWordsWordCount 验证单词数量是否正确
func TestReverseWordsWordCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "验证单词数量1",
			input: "one two three four",
		},
		{
			name:  "验证单词数量2",
			input: "  a   b   c   ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseWords(tt.input)

			// 统计原字符串中的单词数（去除空格后）
			originalWords := 0
			inWord := false
			for _, char := range tt.input {
				if char != ' ' {
					if !inWord {
						originalWords++
						inWord = true
					}
				} else {
					inWord = false
				}
			}

			// 统计结果中的单词数
			resultWords := 0
			inWord = false
			for _, char := range result {
				if char != ' ' {
					if !inWord {
						resultWords++
						inWord = true
					}
				} else {
					inWord = false
				}
			}

			if originalWords != resultWords {
				t.Errorf("reverseWords() 单词数量不匹配: 原字符串 %d 个单词, 结果 %d 个单词", originalWords, resultWords)
			}
		})
	}
}

// BenchmarkSolution 性能测试
func BenchmarkSolution(b *testing.B) {
	testCases := []string{
		"the sky is blue",
		"  hello world  ",
		"a good   example",
		"the quick brown fox jumps over the lazy dog",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testCases {
			reverseWords(s)
		}
	}
}

// BenchmarkReverse 测试 reverse 函数性能
func BenchmarkReverse(b *testing.B) {
	testBytes := []byte("the quick brown fox jumps over the lazy dog")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testCopy := make([]byte, len(testBytes))
		copy(testCopy, testBytes)
		reverse(testCopy)
	}
}
