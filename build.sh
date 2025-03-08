#!/bin/bash

# 定义常见的系统和架构组合
PLATFORMS=("linux/amd64" "linux/arm64" "windows/amd64" "windows/arm64" "darwin/amd64" "darwin/arm64")

# 定义编译命令函数
build_and_check() {
    if [ $1 = "windows" ]; then
        go build -o ./build/gosocks5_$1_$2.exe
    else 
        go build -o ./build/gosocks5_$1_$2
    fi
    if [ $? -ne 0 ]; then
        echo "编译失败"
        exit 1
    fi
}

# 定义分隔函数并检查
separator() {
    for platform in "${PLATFORMS[@]}"; do
    # 分割平台和架构
        IFS='/' read -r GOOS GOARCH <<< "$platform"
        echo "开始编译 $GOOS $GOARCH"
    done
}



main(){
    # 遍历每个平台组合
    for platform in "${PLATFORMS[@]}"; do
        # 分割平台和架构
        IFS='/' read -r GOOS GOARCH <<< "$platform"

        # 设置环境变量
        export GOOS=$GOOS
        export GOARCH=$GOARCH

        # 构建可执行文件
        build_and_check "$GOOS" "$GOARCH"
        echo "编译完成 $GOOS $GOARCH"

    done

    # 恢复默认环境变量
    unset GOOS
    unset GOARCH
}

# separator
main
