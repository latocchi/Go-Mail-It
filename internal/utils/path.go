package utils

import "os"

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		return info.Mode().IsRegular()
	}

	if os.IsNotExist(err) {
		return false
	}
	return false
}

