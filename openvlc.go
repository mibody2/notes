package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 检查命令行参数数量
	if len(os.Args) < 2 {
		fmt.Println("Usage: program vlc://<URL>")
		return
	}

	// 获取第一个命令行参数
	arg := os.Args[1]

	// 检查参数是否以"vlc://"开头
	if !strings.HasPrefix(arg, "vlc://") {
		fmt.Println("Error: argument must start with 'vlc://'")
		return
	}

	// 移除"vlc://"前缀，获取实际的URL
	url := strings.TrimPrefix(arg, "vlc://")

	// 构建VLC命令行调用
	// 注意：这里假设VLC可执行文件在你的系统路径中。如果不是，请使用完整路径。
	cmd := exec.Command("C:\\Program Files\\VideoLAN\\VLC\\vlc.exe", url)

	// 执行命令
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to start VLC: %s\n", err)
		return
	}

	fmt.Printf("VLC has been started with URL: %s\n", url)
}
