.PHONY: all build dev frontend-install frontend-build clean run

# 默认目标
all: build

# 安装前端依赖
frontend-install:
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

# 构建前端
frontend-build:
	@echo "Building frontend..."
	cd frontend && npm run build

# 开发模式（前端热重载）
dev: frontend-install
	@echo "Starting development server..."
	cd frontend && npm run dev

# 构建整个应用
build: frontend-build
	@echo "Building application..."
	go build -o keyview

# 清理
clean:
	@echo "Cleaning..."
	rm -rf frontend/node_modules
	rm -rf frontend/dist
	rm -rf keyview
	rm -f *.db

# 运行应用
run: build
	@echo "Running application..."
	./keyview

# 格式化代码
fmt:
	@echo "Formatting code..."
	go fmt ./...
	cd frontend && npx prettier --write "src/**/*.{js,vue,css}"

# 测试
test:
	@echo "Running tests..."
	go test ./...
