package fileSystem

import "os"


func Exists(somePath string) bool {
	if _, err := os.Stat(somePath); err == nil {
		return true
	}
	return false
}

