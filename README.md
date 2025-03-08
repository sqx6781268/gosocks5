# gosocks5

`gosocks5` 是一个使用 Go 语言编写的 SOCKS5 服务器实现，支持多种操作系统和架构的编译。

## 项目结构
```
build/
  gosocks5_darwin_amd64
  gosocks5_darwin_arm64
  gosocks5_linux_amd64
  gosocks5_linux_arm64
  gosocks5_windows_amd64.exe
  gosocks5_windows_arm64.exe
build.sh
config.json
go.mod
gosocks5.exe
main.go
```

## 功能特性
- 支持多种操作系统和架构的编译，包括 Linux、Windows 和 macOS。
- 支持用户名和密码认证。
- 支持无认证模式。

## 环境要求
- Go 1.19 或更高版本

## 编译和构建
### 编译脚本
项目提供了一个 `build.sh` 脚本，用于编译不同操作系统和架构的可执行文件。你可以运行以下命令来执行脚本：
```bash
./build.sh
```
脚本将为以下平台和架构生成可执行文件：
- `linux/amd64`
- `linux/arm64`
- `windows/amd64`
- `windows/arm64`
- `darwin/amd64`
- `darwin/arm64`

### 手动编译
如果你想手动编译项目，可以使用以下命令：
```bash
# 示例：编译 Linux amd64 版本
GOOS=linux GOARCH=amd64 go build -o ./build/gosocks5_linux_amd64
```

## 配置文件
项目使用 `config.json` 作为配置文件，示例内容如下：
```json
{
  "Port": "1080",
  "User": "user",
  "Password": "pass"
}
```

## 快速入门

### 安装要求
1. 确保已安装 Go 1.19+ 环境
2. 克隆仓库：`git clone https://github.com/yourname/gosocks5.git`
3. 进入项目目录：`cd gosocks5`

### 5分钟快速启动
```bash
# 编译项目
./build.sh

# 使用默认配置运行（无需认证）
./build/gosocks5_linux_amd64 -p 1080

# 测试连接
curl --socks5-hostname 127.0.0.1:1080 https://www.example.com
```

## 详细配置说明

### 配置方式优先级
1. 命令行参数（最高优先级）
2. 配置文件 `config.json`
3. 环境变量

### 完整参数表
| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|-----|
| -c | string | - | 配置文件路径 |
| -p | int | 1080 | 监听端口 |
| -u | string | - | 认证用户名 |
| -P | string | - | 认证密码 |

### 环境变量
```bash
export GOSOCKS5_PORT=1080
export GOSOCKS5_USER=admin
export GOSOCKS5_PASSWORD=secret
```

## 多平台运行示例

### Windows
```powershell
# 编译Windows版本
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o .\build\gosocks5.exe

# 带认证运行
.\build\gosocks5_windows_amd64.exe -p 1080 -u user -P pass
```

### Linux/macOS
```bash
# 后台运行并记录日志
nohup ./gosocks5_linux_amd64 -c config.json > socks5.log 2>&1 &

# 查看运行状态
ps aux | grep gosocks5
```

## 常见问题排查

### 端口占用问题
```bash
# Linux/macOS 查看端口占用
lsof -i :1080

# Windows 查看端口占用
netstat -ano | findstr :1080
```

### 认证失败
1. 检查config.json文件权限
2. 确认用户名密码包含特殊字符时使用引号包裹
3. 查看日志文件确认认证模块加载状态

### 编译问题
1. 确认GOPATH设置正确
2. 检查go mod依赖：`go mod tidy`
3. 清理旧构建：`rm -rf build/*`

## 命令行参数
- `-c`：指定配置文件路径
- `-p`：指定监听端口，默认为 `1080`
- `-u`：指定用户名
- `-P`：指定密码

## 注意事项
- 请确保你的 Go 环境已经正确配置。
- 编译脚本会在 `build` 目录下生成可执行文件。

## 贡献
如果你有任何建议或发现了问题，请随时提交 issue 或 pull request。

## 许可证
本项目采用 [MIT 许可证](LICENSE)。
```

将上述内容保存为 `README.md` 文件，放置在项目的根目录下。这个 `README` 提供了项目的基本信息、使用方法、编译说明等，方便其他开发者了解和使用你的项目。