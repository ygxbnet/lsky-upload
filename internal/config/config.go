package config

import (
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v3"
	"lsky-upload/internal/log"
	"os"
)

//go:embed config-default.yml
var DEFAULT_CONFIG string

// Parse 从默认配置文件路径中获取配置文件内容，并解析为结构体
func Parse(filePath string) (config Result) {
	initFile(filePath)

	// 读取配置文件
	file, err := os.ReadFile(filePath + "/config.yml")
	if err != nil {
		log.Error("❗读取配置文件错误：", err)
		os.Exit(1)
	}

	// 解析配置文件
	conf := Result{}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		log.Error("❗解析配置文件错误：", err)
		os.Exit(1)
	}
	return conf
}

// initFile 初始化及检测配置文件，如果不存在则创建配置文件
func initFile(filePath string) {
	_, err := os.Stat(filePath + "/config.yml")
	// 如果不存在，则创建文件
	if os.IsNotExist(err) {
		file, err := os.Create(filePath + "/config.yml")
		// 如果创建失败，则打印错误信息并退出程序
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		defer file.Close()

		// 将默认配置写入文件
		_, err = file.Write([]byte(DEFAULT_CONFIG))
		// 如果写入失败，则打印错误信息并退出程序
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		// 打印提示信息
		fmt.Println("未发现配置文件，已创建 ", filePath+"/config.yml")
		fmt.Println("请按照文档修改配置文件（按 Enter 键关闭程序）")

		// 等待用户输入
		fmt.Scanln()
		// 退出程序
		os.Exit(0)
	}
}
