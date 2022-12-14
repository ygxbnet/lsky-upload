package main

import (
	"flag"
	"fmt"
	"lsky-upload/internal/config"
	"lsky-upload/internal/httpapi"
	"lsky-upload/internal/tool"
	"os"
)

// 注意在开发时需要将路径传入，例如：-path C:\YGXB\Project\upload
var path = flag.String("path", "", "程序路径")

var configData config.Config

func init() {
	flag.Parse()
}
func main() {
	var programPath string
	if *path == "" {
		programPath = tool.GetProgramPath()
	} else {
		programPath = *path
	}
	configData = config.Parse(programPath)

	url := flag.Args()
	for _, urlStr := range url {
		var getData []byte

		if urlStr[0:4] == "http" {
			getData = httpapi.GetImageData(urlStr)
		} else {
			data, err := os.ReadFile(urlStr)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			getData = data
		}
		fmt.Println(httpapi.UploadImageToLsky(getData, configData.LskyServer, configData.LskyAuthToken))
	}
}
