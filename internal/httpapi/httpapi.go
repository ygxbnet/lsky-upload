package httpapi

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"mime/multipart"
	"net/http"
)

// UploadImageToLsky 上传图片到Lsky，返回图片URL
func UploadImageToLsky(data io.Reader, imageName string, serverURL string, authToken string) string {

	var bufReader bytes.Buffer

	// 生成form表单
	mpWriter := multipart.NewWriter(&bufReader)
	fw, err := mpWriter.CreateFormFile("file", imageName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	io.Copy(fw, data)
	mpWriter.Close()

	// 请求http
	client := &http.Client{}
	req, err := http.NewRequest("POST", serverURL+"/api/v1/upload", &bufReader)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", mpWriter.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return gjson.Parse(string(body)).Get("data").Get("links").Get("url").String()
}

// GetNetworkImageData 请求URL，获取图片数据，返回数据
func GetNetworkImageData(url string) io.Reader {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 Edg/106.0.1370.52")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	pix, _ := io.ReadAll(res.Body)
	return bytes.NewReader(pix)
}
