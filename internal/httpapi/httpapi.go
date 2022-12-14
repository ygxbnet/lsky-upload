package httpapi

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadImageToLsky 上传图片到Lsky，返回图片URL
func UploadImageToLsky(data []byte, server string, authToken string) string {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	fileWriter, _ := bodyWriter.CreateFormFile("files", "file.txt")

	file, _ := os.Open("file.txt")
	defer file.Close()

	io.Copy(fileWriter, file)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("file", string(data))
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", server, payload)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "multipart/form-data")

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

// GetImageData 请求URL，获取图片数据，返回数据
func GetImageData(url string) []byte {
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}
