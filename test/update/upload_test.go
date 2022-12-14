package update

import (
	"bytes"
	"lsky-upload/internal/config"
	"lsky-upload/internal/httpapi"
	"mime/multipart"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {

	var bufReader bytes.Buffer
	mpWriter := multipart.NewWriter(&bufReader)
	fw, err := mpWriter.CreateFormFile("file", "a.txt")
	if err != nil {
		t.Fatal("Create form file error: ", err)
		return
	}

	//f, err2 := os.Open("../../config.yml")
	//if err2 != nil {
	//	t.Fatal(err2)
	//}
	//_, err = io.Copy(fw, f)
	f, err2 := os.ReadFile("../../config.yml")
	if err2 != nil {
		t.Fatal(err2)
	}
	fw.Write(f)

	mpWriter.WriteField("name", "Trump")

	mpWriter.Close()

	t.Log(bufReader.String())
	t.Log(mpWriter.FormDataContentType())
}

func TestUploadImage(t *testing.T) {

	var fileName = "test.png"
	var filePath = `C:\Users\LXM\Pictures\4456352-8938615d16aa9d71.png`

	data, err2 := os.Open(filePath)
	if err2 != nil {
		t.Fatal(err2)
	}

	var cof = config.Parse(`C:\YGXB\Project\lsky-upload`)
	t.Log(httpapi.UploadImageToLsky(data, fileName, cof.LskyServer, cof.LskyAuthToken))
}
