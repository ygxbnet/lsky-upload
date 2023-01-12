package httpapi

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

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
	req, err := http.NewRequest("POST", serverURL+"/api/v1/upload", &bufReader)
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

// GetNetworkImageData 请求URL，获取图片数据，返回数据
func GetNetworkImageData(url string) (data io.ReadCloser, error error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 Edg/106.0.1370.52")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}
