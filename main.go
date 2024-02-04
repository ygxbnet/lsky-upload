package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"lsky-upload/internal/config"
	"lsky-upload/internal/httpapi"
	"lsky-upload/internal/utils"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var configData config.Result

func init() {
	// 解析传入参数
	flag.Parse()
}

func main() {

	// 解析配置
	configData = config.Parse(utils.GetProgramPath())

	// 得到URL地址
	urls := flag.Args()

	// URL分类上传到图床
	for _, url := range urls {
		var getData io.Reader
		var imageName string

		// 读取图片文件，判断是否为网络图片
		if url[0:4] == "http" {
			// 文件为网络图片
			data, err := httpapi.GetNetworkImageData(url)
			defer data.Close()

			// 解析图片类型
			imageType := "webp"
			buff, err := io.ReadAll(data)
			fileType := http.DetectContentType(buff)
			if fileType[:5] == "image" { // fileType例子：image/jpeg
				imageType = fileType[6:]
			} else {
				fmt.Println("❗输入的网络链接不是图片，请检查链接是否正确\n", string(buff)[:1000])
				os.Exit(1)
			}

			imageName = fmt.Sprintf("%s.%s", time.Now().Format("2006-01-02 15:04:05"), imageType)
			if err != nil {
				fmt.Println("❗获取网络图片错误：", err)
				return
			}
			getData = bytes.NewReader(buff)
		} else {
			imageName = filepath.Base(url)

			data, err := os.Open(url)
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
				fmt.Println(gjson.Parse(string(returnMessage)).Get("data.links.url").String())
			} else {
				// 上传失败
				fmt.Println("❗上传图片失败 服务器返回信息：", string(returnMessage))
			}
		} else {
			// 请求上传图片接口失败
			fmt.Printf("❗请求接口%d \t\n详细信息：%s", response.StatusCode, string(returnMessage))
		}
	}
}
