package main

import (
	"flag"
	"fmt"
	"io"
	"lsky-upload/internal/config"
	"lsky-upload/internal/httpapi"
	"lsky-upload/internal/tool"
	"os"
	"path/filepath"
	"time"
)

// 注意在开发时需要将路径传入，例如：-PATH C:\YGXB\Project\upload
var PATH = flag.String("PATH", "", "程序路径")

var configData config.Config

func init() {
	flag.Parse()
}
func main() {
	// 解析配置文件
	var programPath string
	if *PATH == "" {
		path, err := tool.GetProgramPath()
		if err != nil {
			fmt.Println("获取程序路径错误：", err)
			os.Exit(1)
		}
		programPath = path
	} else {
		programPath = *PATH
	}
	configData = config.Parse(programPath)

	// 得到URL地址
	url := flag.Args()

	// URL分类上传到图床
	for _, urlString := range url {
		var getData io.Reader
		var imageName string

		if urlString[0:4] == "http" {
			imageName = fmt.Sprintf("%s.webp", time.Now().Format("2006-01-02 15:04:05"))
			data, err := httpapi.GetNetworkImageData(urlString)
			if err != nil {
				fmt.Println("获取网络图片错误：", err)
				return
			}
			getData = data
		} else {
			imageName = filepath.Base(urlString)

			data, err := os.Open(urlString)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			getData = data
		}

		imageURL, err := httpapi.UploadImageToLsky(getData, imageName, configData.LskyServer, configData.LskyAuthToken)
		if err != nil {
			fmt.Println("上传图片错误：", err)
			return
		}
		fmt.Println(imageURL.Get("data").Get("links").Get("url").String())
	}
}
