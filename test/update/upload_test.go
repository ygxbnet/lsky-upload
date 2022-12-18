package update

import (
	"bytes"
	"lsky-upload/internal/config"
	"lsky-upload/internal/httpapi"
	"mime/multipart"
	"os"
	"testing"
)

func Test_Upload(t *testing.T) {
	var bufReader bytes.Buffer
	mpWriter := multipart.NewWriter(&bufReader)
	fw, err := mpWriter.CreateFormFile("file", "a.txt")
	if err != nil {
		t.Fatal("Create form file error: ", err)
		return
	}

	// f, err2 := os.Open("../../config.yml")
	// if err2 != nil {
	//	t.Fatal(err2)
	// }
	// _, err = io.Copy(fw, f)
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

func Test_Upload_Local_Image(t *testing.T) {
	var fileName = "test.png"
	var filePath = `C:\Users\LXM\Pictures\4456352-8938615d16aa9d71.png`

	data, err2 := os.Open(filePath)
	if err2 != nil {
		t.Fatal(err2)
	}

	var cof = config.Parse(`C:\YGXB\Project\lsky-upload`)
	t.Log(httpapi.UploadImageToLsky(data, fileName, cof.LskyServer, cof.LskyAuthToken))
}

func Test_Upload_Network_Image(t *testing.T) {
	var fileName = "test.webp"
	var fileURL = `https://ts1.cn.mm.bing.net/th/id/R-C.c2f6165e1acf6b986a20e8b676f13bdd?rik=Dm%2bD%2fgeM4oqgBw&riu=http%3a%2f%2fwww.quazero.com%2fuploads%2fallimg%2f140303%2f1-140303215045.jpg&ehk=sdIsCQj%2bvjKfUs%2fDw%2fZekILroLb1ALwbKghApSPIq4U%3d&risl=&pid=ImgRaw&r=0`
	data := httpapi.GetNetworkImageData(fileURL)

	var cof = config.Parse(`C:\YGXB\Project\lsky-upload`)
	t.Log(httpapi.UploadImageToLsky(data, fileName, cof.LskyServer, cof.LskyAuthToken))
}
