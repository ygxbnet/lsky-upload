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
var PATH = *flag.String("PATH", "", "程序路径")

var config_data config.Config

func init() {
	flag.Parse()
}
func main() {
	// 解析配置文件
	var programPath string
	if PATH == "" {
		programPath = tool.GetProgramPath()
	} else {
		programPath = PATH
	}
	config_data = config.Parse(programPath)

	// 得到URL地址
	url := flag.Args()

	// URL分类上传到图床
	for _, urlString := range url {
		var getData io.Reader
		var imageName string

		if urlString[0:4] == "http" {
			imageName = fmt.Sprintf("%s.webp", time.Now().Format("2006-01-02 15:04:05"))
			getData = httpapi.GetNetworkImageData(urlString)
		} else {
			imageName = filepath.Base(urlString)

			data, err := os.Open(urlString)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			getData = data
		}

		imageURL := httpapi.UploadImageToLsky(getData, imageName, config_data.LskyServer, config_data.LskyAuthToken)
		if imageURL != "" {
			fmt.Println(imageURL)
		}
	}
}
