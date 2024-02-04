package utils

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// PATH 注意在开发时需要将路径传入
//
// 例如：-path C:\xxx\xxx
var PATH = flag.String("path", "", "程序路径")

// getProgramPath 获取程序路径
func getProgramPath() (programPath string, error error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}

// GetProgramPath 获取程序路径外部封装
func GetProgramPath() (programPath string) {
	if *PATH == "" {
		path, err := getProgramPath()
		if err != nil {
			fmt.Println("❗无法获取当前程序路径：", err)
			os.Exit(1)
		}
		return path
	} else {
		// 开发时运行路径与项目路径不同,需使用手动传入的 config.yml 地址
		return *PATH
	}
}
