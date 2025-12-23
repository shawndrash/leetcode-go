.PHONY: test test-all test-verbose bench clean help

# 默认目标
help:
	@echo "LeetCode Go 项目 Makefile"
	@echo ""
	@echo "使用方法:"
	@echo "  make test NUM=1          - 测试指定题目（例如：第1题）"
	@echo "  make test-all            - 测试所有题目"
	@echo "  make test-verbose NUM=1  - 详细输出测试指定题目"
	@echo "  make bench NUM=1         - 性能测试指定题目"
	@echo "  make bench-all           - 性能测试所有题目"
	@echo "  make clean               - 清理测试缓存"
	@echo "  make new NUM=1 NAME=xxx  - 创建新题目（需要 CNAME 参数）"
	@echo ""
	@echo "示例:"
	@echo "  make test NUM=1"
	@echo "  make bench NUM=2"
	@echo "  make new NUM=3 NAME=longest-substring CNAME=\"无重复字符的最长子串\""

# 测试单个题目
test:
ifndef NUM
	@echo "错误: 请指定题目编号"
	@echo "用法: make test NUM=1"
	@exit 1
endif
	@PADDED=$$(printf "%04d" $(NUM)); \
	DIR=$$(ls -d problems/$$PADDED.* 2>/dev/null | head -n 1); \
	if [ -z "$$DIR" ]; then \
		echo "错误: 找不到第 $(NUM) 题"; \
		exit 1; \
	fi; \
	echo "测试: $$DIR"; \
	cd $$DIR && go test

# 详细测试单个题目
test-verbose:
ifndef NUM
	@echo "错误: 请指定题目编号"
	@echo "用法: make test-verbose NUM=1"
	@exit 1
endif
	@PADDED=$$(printf "%04d" $(NUM)); \
	DIR=$$(ls -d problems/$$PADDED.* 2>/dev/null | head -n 1); \
	if [ -z "$$DIR" ]; then \
		echo "错误: 找不到第 $(NUM) 题"; \
		exit 1; \
	fi; \
	echo "详细测试: $$DIR"; \
	cd $$DIR && go test -v

# 测试所有题目
test-all:
	@echo "测试所有题目..."
	@go test ./problems/...

# 性能测试单个题目
bench:
ifndef NUM
	@echo "错误: 请指定题目编号"
	@echo "用法: make bench NUM=1"
	@exit 1
endif
	@PADDED=$$(printf "%04d" $(NUM)); \
	DIR=$$(ls -d problems/$$PADDED.* 2>/dev/null | head -n 1); \
	if [ -z "$$DIR" ]; then \
		echo "错误: 找不到第 $(NUM) 题"; \
		exit 1; \
	fi; \
	echo "性能测试: $$DIR"; \
	cd $$DIR && go test -bench=. -benchmem

# 性能测试所有题目
bench-all:
	@echo "性能测试所有题目..."
	@go test -bench=. -benchmem ./problems/...

# 测试覆盖率
coverage:
ifndef NUM
	@echo "所有题目的测试覆盖率..."
	@go test -cover ./problems/...
else
	@PADDED=$$(printf "%04d" $(NUM)); \
	DIR=$$(ls -d problems/$$PADDED.* 2>/dev/null | head -n 1); \
	if [ -z "$$DIR" ]; then \
		echo "错误: 找不到第 $(NUM) 题"; \
		exit 1; \
	fi; \
	echo "测试覆盖率: $$DIR"; \
	cd $$DIR && go test -cover
endif

# 创建新题目
new:
ifndef NUM
	@echo "错误: 请指定题目编号"
	@echo "用法: make new NUM=3 NAME=longest-substring CNAME=\"无重复字符的最长子串\""
	@exit 1
endif
ifndef NAME
	@echo "错误: 请指定题目名称（英文）"
	@echo "用法: make new NUM=3 NAME=longest-substring CNAME=\"无重复字符的最长子串\""
	@exit 1
endif
	@./new_problem.sh $(NUM) $(NAME) "$(CNAME)"

# 清理测试缓存
clean:
	@echo "清理测试缓存..."
	@go clean -testcache

# 列出所有题目
list:
	@echo "已完成的题目:"
	@ls -1 problems/ | nl

# 格式化代码
fmt:
	@echo "格式化代码..."
	@go fmt ./...

# 代码检查
lint:
	@echo "代码检查..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "未安装 golangci-lint，跳过检查"; \
		echo "安装方法: brew install golangci-lint"; \
	fi

