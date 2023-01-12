package get_file_type

import (
	"net/http"
	"testing"
)

func TestGetFileType(t *testing.T) {
	var fileURL = `https://ts1.cn.mm.bing.net/th/id/R-C.c2f6165e1acf6b986a20e8b676f13bdd?rik=Dm%2bD%2fgeM4oqgBw&riu=http%3a%2f%2fwww.quazero.com%2fuploads%2fallimg%2f140303%2f1-140303215045.jpg&ehk=sdIsCQj%2bvjKfUs%2fDw%2fZekILroLb1ALwbKghApSPIq4U%3d&risl=&pid=ImgRaw&r=0`
	data, err := http.Get(fileURL)
	if err != nil {
		t.Error(err)
	}

	buff := make([]byte, 512) // 只读取512字节，减少内存开销
	_, err = data.Body.Read(buff)

	fileType := http.DetectContentType(buff)
	t.Log(fileType)
	if fileType[:5] == "image" {
		t.Log(fileType[6:])
	}
}
