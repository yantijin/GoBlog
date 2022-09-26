package utils

import (
	"os"
)

// 确定文件是否存在
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true, nil
		}
		return false, err
	}
	return false, err
}
