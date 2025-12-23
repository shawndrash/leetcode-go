#!/bin/bash

# 使用方法: ./new_problem.sh <题号> <题目名称> [中文名称]
# 示例: ./new_problem.sh 1 two-sum "两数之和"

if [ $# -lt 2 ]; then
    echo "使用方法: ./new_problem.sh <题号> <题目名称> [中文名称]"
    echo "示例: ./new_problem.sh 1 two-sum \"两数之和\""
    exit 1
fi

PROBLEM_NUM=$1
PROBLEM_NAME=$2
CHINESE_NAME=${3:-""}

# 格式化题号为4位数字
FORMATTED_NUM=$(printf "%04d" $PROBLEM_NUM)

# 创建目录名（例如：0001.two-sum）
DIR_NAME="${FORMATTED_NUM}.${PROBLEM_NAME}"
PROBLEM_DIR="problems/${DIR_NAME}"

# 包名（将横线转换为下划线）
PACKAGE_NAME=$(echo $PROBLEM_NAME | tr '-' '_')

# 函数名（驼峰命名）
FUNCTION_NAME=$(echo $PROBLEM_NAME | sed -E 's/-([a-z])/\U\1/g')

echo "正在创建题目: ${DIR_NAME}"

# 创建目录
mkdir -p "${PROBLEM_DIR}"

# 创建 README.md
cat > "${PROBLEM_DIR}/README.md" <<EOF
# ${FORMATTED_NUM}. ${PROBLEM_NAME}

**难度**: \`Medium\`

## 题目描述

${CHINESE_NAME}

TODO: 添加题目描述

## 示例

\`\`\`
示例 1:
输入:
输出:
解释:
\`\`\`

## 约束条件

- TODO: 添加约束条件

## 解题思路

### 方法一：

TODO: 添加解题思路

**时间复杂度**: O(n)
**空间复杂度**: O(1)

## 相关题目

- TODO: 添加相关题目链接
EOF

# 创建 solution.go
cat > "${PROBLEM_DIR}/solution.go" <<EOF
package ${PACKAGE_NAME}

// ${FUNCTION_NAME} 解决 LeetCode 第 ${PROBLEM_NUM} 题
func ${FUNCTION_NAME}() {
	// TODO: 实现你的解法
}
EOF

# 创建 solution_test.go
cat > "${PROBLEM_DIR}/solution_test.go" <<EOF
package ${PACKAGE_NAME}

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
			// got := ${FUNCTION_NAME}(...)
			// if got != tt.want {
			//     t.Errorf("${FUNCTION_NAME}() = %v, want %v", got, tt.want)
			// }
		})
	}
}

// BenchmarkSolution 性能测试
func BenchmarkSolution(b *testing.B) {
	// TODO: 添加性能测试
	for i := 0; i < b.N; i++ {
		// ${FUNCTION_NAME}(...)
	}
}
EOF

echo "✅ 题目目录创建成功: ${PROBLEM_DIR}"
echo ""
echo "下一步:"
echo "  1. 编辑 ${PROBLEM_DIR}/solution.go 实现解法"
echo "  2. 编辑 ${PROBLEM_DIR}/solution_test.go 添加测试用例"
echo "  3. 运行测试: cd ${PROBLEM_DIR} && go test -v"
echo ""

