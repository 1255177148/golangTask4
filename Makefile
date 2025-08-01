APP_NAME := myapp
MAIN := ./cmd/main.go
BIN_DIR := ./bin
BIN_FILE := $(BIN_DIR)/$(APP_NAME)

# 默认环境(dev/test/prod)
ENV ?= dev

# 设置 APP_ENV 环境变量
export APP_ENV=$(ENV)

.DEFAULT_GOAL := run

# 初始化依赖
init:
	go mod tidy
	go mod download

# 运行项目
run:
	@echo "🚀 运行环境: $(ENV)"
	go run $(MAIN)

# 构建二进制
build:
	@echo "🔨 构建 $(APP_NAME) 环境: $(ENV)"
	mkdir -p $(BIN_DIR)
	APP_ENV=$(ENV) go build -o $(BIN_FILE) $(MAIN)

# 清理构建产物
clean:
	rm -rf $(BIN_DIR)

# 运行测试
test:
	APP_ENV=$(ENV) go test ./... -v

# 整理依赖
tidy:
	go mod tidy

# 代码检查
lint:
	go vet ./...
