package config

import (
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

//go:embed config-default.yml
var defaultConfig string

// Parse 从默认配置文件路径中获取
func Parse(filePath string) Config {

	initFile(filePath)

	file, err := os.ReadFile(filePath + "/config.yml")
	if err != nil {
		fmt.Println("读取配置文件错误", err)
		os.Exit(1)
	}

	config := Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("解析配置文件错误", err)
		os.Exit(1)
	}
	return config
}

// initFile 初始化及检测配置文件，如果不存在则创建配置文件
func initFile(filePath string) {
	
	_, err := os.Stat(filePath + "/config.yml")
	if os.IsNotExist(err) {
		file, err := os.Create(filePath + "/config.yml")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		_, err = file.Write([]byte(defaultConfig))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("未发现配置文件，已创建 ", filePath+"/config.yml")
		fmt.Println("请修改配置文件后再重新启动")
		os.Exit(0)
	}
}
