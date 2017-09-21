package _resources

import (
	"testing"
	"github.com/powergun/goUsage/fileSystem"
)


func TestResourcePathGetter(t *testing.T) {
	filePath := ResourcePath("readme.txt")
	if !fileSystem.Exists(filePath) {
		t.Error("fail to get resource file: readme.txt")
	}
	filePath = ResourcePath("nothere.ext")
	if fileSystem.Exists(filePath) {
		t.Error("not expected: false positive")
	}
	dirPath := ResourcePath("conf")
	if !fileSystem.Exists(dirPath) {
		t.Error("not expected")
	}
}

