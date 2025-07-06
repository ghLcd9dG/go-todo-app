# Todo CLI - 命令行待办事项管理工具

一个简单而强大的命令行待办事项管理工具，使用Go语言开发。

## 功能特性

- ✅ 添加待办事项
- 📋 列出所有待办事项
- ✔️ 标记待办事项为完成
- 💾 自动保存到JSON文件
- ⏰ 显示创建时间

## 安装

### 方法1：直接安装（推荐）

```bash
go install github.com/af-liu/todo-cli@latest
```

安装完成后，确保您的`$GOPATH/bin`或`$GOBIN`目录在您的PATH环境变量中。

### 方法2：从源码安装

1. 克隆仓库：
```bash
git clone https://github.com/af-liu/todo-cli.git
cd todo-cli
```

2. 编译并安装：
```bash
go install
```

### 方法3：本地安装

如果您有源码，可以在项目目录中运行：
```bash
go install .
```

## 使用方法

安装后，您可以在命令行的任何位置使用`todo-cli`命令：

```bash
todo-cli
```

### 可用命令

- `add <任务描述>` - 添加一个新的待办事项
- `list` - 列出所有待办事项
- `complete <ID>` - 将指定ID的待办事项标记为完成
- `help` - 显示帮助信息
- `exit` - 退出程序

### 使用示例

```bash
# 启动程序
todo-cli

# 在交互模式下：
> add 学习Go语言
> add 完成项目文档
> list
> complete 1
> exit
```

## 数据存储

待办事项数据保存在当前工作目录的`todos.json`文件中。

## 系统要求

- Go 1.24 或更高版本

## 许可证

MIT License
