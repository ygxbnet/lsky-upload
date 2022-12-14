package tool

import (
	"log"
	"os"
	"path/filepath"
)

func GetProgramPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return dir
}
