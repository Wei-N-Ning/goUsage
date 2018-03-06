package fileIO

import (
	"runtime"
	"path/filepath"
	"os"
	"strings"
)


func resourcePath(partialPath string) string {
	_, filePath, _, _ := runtime.Caller(1)
	tokens := []string{filepath.Dir(filePath), "resources", partialPath}
	return  strings.Join(tokens, string(os.PathSeparator))
}