.PHONY: all build dev frontend-install frontend-build clean run mac-app dev-mac

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

# 开发模式 - 运行 Wails 开发服务器
dev: frontend-install
	@echo "Starting Wails development mode..."
	wails3 dev

# 构建整个应用
build: frontend-build
	@echo "Building application..."
	go build -o keyview

# 构建 macOS 应用
mac-app:
	@echo "Building macOS application..."
	wails3 package
	@echo "✅ macOS app built successfully!"
	@echo "📍 Location: bin/KeyView.app"
	@echo ""
	@echo "Run with: make run-mac or open bin/KeyView.app"


# 运行应用
run: build
	@echo "Running application..."
	./keyview

# 运行构建好的 macOS 应用
run-mac: mac-app
	@echo "Running macOS application..."
	open bin/KeyView.app


# 格式化代码
fmt:
	@echo "Formatting code..."
	go fmt ./...
	cd frontend && npx prettier --write "src/**/*.{js,vue,css}"

# 测试
test:
	@echo "Running tests..."
	go test ./...
