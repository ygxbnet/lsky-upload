package httpapi

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0"
var apiPath = "/api/v1/upload"

// UploadImageToLsky 上传图片到Lsky，返回相关信息
func UploadImageToLsky(data io.Reader, imageName string, serverURL string, authToken string) (response *http.Response, error error) {
	var bufReader bytes.Buffer

	// 生成form表单
	mpWriter := multipart.NewWriter(&bufReader)
	fw, err := mpWriter.CreateFormFile("file", imageName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, data)
	if err != nil {
		return nil, err
	}
	mpWriter.Close()

	// 请求http
	client := &http.Client{}
	req, err := http.NewRequest("POST", serverURL+apiPath, &bufReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", mpWriter.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetNetworkImageData 获取网络图片数据
func GetNetworkImageData(url string) (data io.ReadCloser, error error) {
	// 创建一个http.Client实例
	client := &http.Client{}
	// 创建一个GET请求
	req, err := http.NewRequest("GET", url, nil)
	// 如果创建请求失败，返回错误
	if err != nil {
		return nil, err
	}
	// 设置请求头中的User-Agent
	req.Header.Add("User-Agent", userAgent)

	// 执行请求
	res, err := client.Do(req)
	// 如果执行请求失败，返回错误
	if err != nil {
		return nil, err
	}

	// 返回图片数据和nil
	return res.Body, nil
}
