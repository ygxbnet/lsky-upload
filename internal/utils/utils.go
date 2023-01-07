package utils

import (
	"os"
	"path/filepath"
)

// GetProgramPath 获取程序路径
func GetProgramPath() (programPath string, error error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}
