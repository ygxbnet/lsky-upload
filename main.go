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

// 注意在开发时需要将路径传入，例如：-path C:\YGXB\Project\upload
var path = flag.String("path", "", "程序路径")

var configData config.Config

func init() {
	flag.Parse()
}
func main() {
	// 解析配置文件
	var programPath string
	if *path == "" {
		programPath = tool.GetProgramPath()
	} else {
		programPath = *path
	}
	configData = config.Parse(programPath)

	// 得到URL地址
	url := flag.Args()

	// URL分类上传到图床
	for _, urlStr := range url {
		var getData io.Reader
		var imageName string

		if urlStr[0:4] == "http" {
			imageName = fmt.Sprintf("%s.webp", time.Now().Format("2006-01-02 15:04:05"))
			getData = httpapi.GetNetworkImageData(urlStr)
		} else {
			imageName = filepath.Base(urlStr)

			data, err := os.Open(urlStr)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			getData = data
		}

		imageURL := httpapi.UploadImageToLsky(getData, imageName, configData.LskyServer, configData.LskyAuthToken)
		if imageURL != "" {
			fmt.Println(imageURL)
		}
	}
}
