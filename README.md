# LeetCode Go 题解集

一个用 Go 语言刷 LeetCode 的项目模板，支持本地测试和题目管理。

## 项目结构

```
leetcode-go/
├── go.mod                    # Go 模块文件
├── README.md                 # 项目说明
├── new_problem.sh           # 快速创建新题目的脚本
├── problems/                 # 所有题目目录
│   ├── 0001.two-sum/        # 题目目录（格式：题号.题目名）
│   │   ├── solution.go      # 解法实现
│   │   ├── solution_test.go # 测试用例
│   │   └── README.md        # 题目描述
│   └── ...
└── utils/                    # 通用工具和数据结构
    ├── listnode.go          # 链表节点
    ├── treenode.go          # 树节点
    └── helpers.go           # 辅助函数
```

## 快速开始

### 1. 创建新题目

使用脚本快速创建题目模板：

```bash
./new_problem.sh 1 two-sum "两数之和"
```

这会在 `problems/0001.two-sum/` 目录下创建模板文件。

### 2. 编写解法

在 `solution.go` 中实现你的解法：

```go
package two_sum

func twoSum(nums []int, target int) []int {
    // 你的解法
    return nil
}
```

### 3. 编写测试用例

在 `solution_test.go` 中添加测试：

```go
func TestTwoSum(t *testing.T) {
    tests := []struct {
        name   string
        nums   []int
        target int
        want   []int
    }{
        {"示例1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := twoSum(tt.nums, tt.target)
            // 断言检查
        })
    }
}
```

### 4. 运行测试

```bash
# 测试单个题目
cd problems/0001.two-sum
go test -v

# 测试所有题目
go test ./problems/...

# 运行性能测试
go test -bench=. ./problems/0001.two-sum
```

## 命名规范

- **目录名**：`0001.two-sum`（题号.题目slug）
- **包名**：`two_sum`（使用下划线）
- **函数名**：与 LeetCode 保持一致

## 特性

✅ 每道题独立包，互不冲突
✅ 完整的单元测试支持
✅ 性能基准测试
✅ 通用数据结构（链表、树等）
✅ 快速生成题目模板

## 常用命令

```bash
# 运行所有测试
go test ./...

# 查看测试覆盖率
go test -cover ./problems/...

# 运行带详细输出的测试
go test -v ./problems/0001.two-sum

# 运行性能测试
go test -bench=. -benchmem ./problems/...
```

## 目录说明

- `problems/`: 存放所有 LeetCode 题目，每题一个独立目录
- `utils/`: 存放常用的数据结构和辅助函数
- `new_problem.sh`: 自动化创建新题目目录结构的脚本

## 贡献

欢迎提交 PR 添加更多题解！

## License

MIT

