package log

import (
	"fmt"
	"log"
	"lsky-upload/internal/utils"
	"os"
	"path"
)

func Error(a ...any) {
	writeToFile(a)
	fmt.Println(a...)
}

func writeToFile(a ...any) {
	f, err := os.OpenFile(path.Join(utils.GetProgramPath(), "lsky-upload.log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("❗无法将错误日志写入 lsky-upload.log", err)
		return
	}
	defer f.Close()

	log.SetOutput(f)                    // 设置log输出
	log.SetFlags(log.Ldate | log.Ltime) // 设置log前缀
	log.Println("\n", a)
	fmt.Println("错误日志已写入到", path.Join(utils.GetProgramPath(), "lsky-upload.log"))
}
