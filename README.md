# KeyView - 键盘使用历史记录查看工具

基于 Wails v3.0.0-alpha.51 + Vue3 + github.com/wangle201210/keylogger + SQLite3 实现的键盘使用历史记录查看工具。

## 功能特性

- 🔤 **实时键盘监听** - 记录所有键盘按键事件
- 💾 **SQLite 数据存储** - 持久化存储键盘历史记录
- 🔍 **强大的筛选功能** - 按按键名称、日期、动作类型筛选
- 📊 **统计分析** - 显���总记录数、今日按键次数等统计信息
- 🎨 **现代化 UI** - 基于 Vue3 构建的渐变风格界面
- ⚡ **高性能** - 支持分页显示，流畅浏览大量记录

## 技术栈

### 后端
- **Go 1.24+** - 主要编程语言
- **Wails v3.0.0-alpha.51** - 桌面应用框架
- **github.com/wangle201210/keylogger** - 键盘事件监听库
- **GORM + SQLite3** - 数据库 ORM

### 前端
- **Vue 3** - 渐进式 JavaScript 框架
- **Vite** - 快速的前端构建工具

## 项目结构

```
keyview/
├── main.go              # 应用入口
├── internal/
│   └── app/
│       └── service.go   # 后端服务实现
├── frontend/            # Vue3 前端
│   ├── src/
│   │   ├── App.vue     # 主应用组件
│   │   ├── main.js     # 前端入口
│   │   ├── style.css   # 全局样式
│   │   └── services/
│   │       └── wails.js # Wails API 封装
│   ├── index.html
│   ├── package.json
│   └── vite.config.js
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 安装与运行

### 前置要求

- Go 1.24 或更高版本
- Node.js 18+ 和 npm
- macOS (目前仅支持 macOS 平台)
- Xcode 命令行工具 (用于编译 C 代码)

### 安装依赖

1. 克隆项目（包括 keylogger 子模块）:
```bash
git clone --recurse-submodules <repository-url>
cd keyview
```

2. 安装前端依赖:
```bash
make frontend-install
# 或
cd frontend && npm install
```

3. 下载 Go 依赖:
```bash
go mod tidy
```

### 运行应用

#### 方式一：使用 Makefile

```bash
# 构建前端并运行应用
make run

# 或者分步执行
make frontend-build  # 构建前端
make build          # 构建应用
./keyview          # 运行
```

#### 方式二：手动执行

```bash
# 1. 构建前端
cd frontend
npm run build
cd ..

# 2. 构建并运行应用
go build -o keyview
./keyview
```

### 开发模式

仅开发前端（热重载）:
```bash
cd frontend
npm run dev
```

前端开发服务器将在 `http://localhost:5173` 启动。

## 使用说明

1. **启动应用** - 运行应用后会自动打开窗口
2. **开始记录** - 点击"开始记录"按钮开始监听键盘事件
3. **查看记录** - 所有按键事件会实时显示在列表中
4. **筛选数据** - 使用顶部筛选栏按条件过滤记录
5. **停止记录** - 点击"停止记录"按钮暂停监听

### 筛选功能

- **按键筛选** - 选择特定的按键名称查看其历史记录
- **日期筛选** - 按日期查看特定日期的记录
- **动作筛选** - 选择"按下"或"释放"查看特定动作

### 数据统计

- **总记录数** - 数据库中存储的总记录数量
- **今日按键** - 今天按下的按键次数（仅计算按下事件）

## 隐私说明

⚠️ **重要提示**：
- 本工具仅用于个人使用统计和分析目的
- 键盘记录数据存储在本地 SQLite 数据库中
- 请勿将此工具用于恶意目的或未经授权的监控
- 建议定期清理旧的记录数据

## 数据存储位置

数据库文件默认存储在应用运行目录下的 `keyview.db` 文件中。

## 权限要求

在 macOS 上运行键盘监听需要辅助功能权限：
1. 打开"系统设置" > "隐私与安全性" > "辅助功能"
2. 添加 `keyview` 应用到允许列表
3. 重启应用

## 故障排除

### 构建失败

如果遇到 CGO 相关错误，确保已安装 Xcode 命令行工具：
```bash
xcode-select --install
```

### 权限被拒绝

如果在启动时遇到权限错误，请按照上述步骤授予辅助功能权限。

### 前端无法加载

确保前端已经构建：
```bash
cd frontend && npm run build
```

## 开发

### 代码格式化

```bash
make fmt
```

### 清理构建文件

```bash
make clean
```

## 许可证

请参考项目的 LICENSE 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 致谢

- [Wails](https://wails.io/) - 优秀的 Go 桌面应用框架
- [Vue.js](https://vuejs.org/) - 渐进式 JavaScript 框架
- [GORM](https://gorm.io/) - Go 的 ORM 库
