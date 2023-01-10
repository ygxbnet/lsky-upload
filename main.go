package main

import (
	"flag"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"lsky-upload/internal/config"
	"lsky-upload/internal/httpapi"
	"lsky-upload/internal/utils"
	"os"
	"path/filepath"
	"time"
)

// PATH 注意在开发时需要将路径传入，例如：-path C:\YGXB\Project\upload
var PATH = flag.String("path", "", "程序路径")

var configData config.Result

func init() {
	// 解析传入参数
	flag.Parse()
}

func main() {
	// 解析配置文件
	var programPath string
	if *PATH == "" {
		path, err := utils.GetProgramPath()
		if err != nil {
			fmt.Println("❗无法获取当前程序路径：", err)
			os.Exit(1)
		}
		programPath = path
	} else {
		programPath = *PATH
	}
	configData = config.Parse(programPath)

	// 得到URL地址
	urls := flag.Args()

	// URL分类上传到图床
	for _, urlString := range urls {
		var getData io.Reader
		var imageName string

		// 读取图片文件
		if urlString[0:4] == "http" {
			imageName = fmt.Sprintf("%s.webp", time.Now().Format("2006-01-02 15:04:05"))
			data, err := httpapi.GetNetworkImageData(urlString)
			if err != nil {
				fmt.Println("❗获取网络图片错误：", err)
				return
			}
			getData = data
		} else {
			imageName = filepath.Base(urlString)

			data, err := os.Open(urlString)
			if err != nil {
				fmt.Println("❗打开文件失败", err)
				os.Exit(1)
			}
			getData = data
		}

		// 上传图片到图床
		response, err := httpapi.UploadImageToLsky(getData, imageName, configData.LskyServer, configData.LskyAuthToken)
		if err != nil {
			fmt.Println("❗上传图片错误：", err)
			return
		}
		defer response.Body.Close()

		returnMessage, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("❗读取http返回信息失败：", err)
		}

		// 处理返回结果
		if response.StatusCode == 200 {
			if gjson.Parse(string(returnMessage)).Get("status").String() == "true" {
				// 成功上传
				fmt.Println(gjson.Parse(string(returnMessage)).Get("data").Get("links").Get("url").String())
			} else {
				// 上传失败
				fmt.Println("❗上传图片失败 服务器返回信息：", gjson.Parse(string(returnMessage)).Get("message").String())
			}
		} else {
			// 请求上传图片接口失败
			fmt.Printf("❗请求接口%d \t\n详细信息：%s", response.StatusCode, string(returnMessage))
		}
	}
}
