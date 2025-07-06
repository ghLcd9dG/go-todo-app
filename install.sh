#!/bin/bash

# Todo CLI 安装脚本

echo "正在安装 Todo CLI..."

# 检查Go是否已安装
if ! command -v go &> /dev/null; then
    echo "错误: 未找到Go。请先安装Go语言环境。"
    echo "访问 https://golang.org/dl/ 下载并安装Go"
    exit 1
fi

# 显示Go版本
echo "检测到Go版本: $(go version)"

# 安装程序
echo "正在编译并安装..."
go install .

if [ $? -eq 0 ]; then
    echo "✅ Todo CLI 安装成功！"
    echo ""
    echo "使用方法："
    echo "  在命令行中输入 'todo-cli' 即可开始使用"
    echo ""
    echo "注意："
    echo "  请确保您的 \$GOPATH/bin 或 \$GOBIN 目录在 PATH 环境变量中"
    echo "  如果命令找不到，请运行："
    echo "    export PATH=\$PATH:\$(go env GOPATH)/bin"
else
    echo "❌ 安装失败，请检查错误信息"
    exit 1
fi
