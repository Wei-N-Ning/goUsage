package _resources

import (
	"runtime"
	"path/filepath"
	"os"
)


func ResourcePath(partialPath string) string {
	_, filePath, _, _ := runtime.Caller(1)
	return filepath.Dir(filePath) + string(os.PathSeparator) + partialPath
}
