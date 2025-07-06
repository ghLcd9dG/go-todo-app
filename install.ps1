# Todo CLI 安装脚本 (PowerShell)

Write-Host "正在安装 Todo CLI..." -ForegroundColor Green

# 检查Go是否已安装
try {
    $goVersion = go version
    Write-Host "检测到Go版本: $goVersion" -ForegroundColor Blue
} catch {
    Write-Host "错误: 未找到Go。请先安装Go语言环境。" -ForegroundColor Red
    Write-Host "访问 https://golang.org/dl/ 下载并安装Go" -ForegroundColor Yellow
    exit 1
}

# 安装程序
Write-Host "正在编译并安装..." -ForegroundColor Yellow
go install .

if ($LASTEXITCODE -eq 0) {
    Write-Host "✅ Todo CLI 安装成功！" -ForegroundColor Green
    Write-Host ""
    Write-Host "使用方法：" -ForegroundColor Cyan
    Write-Host "  在命令行中输入 'todo-cli' 即可开始使用"
    Write-Host ""
    Write-Host "注意：" -ForegroundColor Yellow
    Write-Host "  请确保您的 GOPATH\bin 或 GOBIN 目录在 PATH 环境变量中"
    Write-Host "  如果命令找不到，请将以下路径添加到系统PATH："
    $goPath = go env GOPATH
    Write-Host "    $goPath\bin" -ForegroundColor Magenta
} else {
    Write-Host "❌ 安装失败，请检查错误信息" -ForegroundColor Red
    exit 1
}
