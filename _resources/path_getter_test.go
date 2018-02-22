package _resources

import (
	"testing"
	"github.com/powergun/goUsage/file_system"
)


func TestResourcePathGetter(t *testing.T) {
	filePath := ResourcePath("readme.txt")
	if !file_system.Exists(filePath) {
		t.Error("fail to get resource file: readme.txt")
	}
	filePath = ResourcePath("nothere.ext")
	if file_system.Exists(filePath) {
		t.Error("not expected: false positive")
	}
	dirPath := ResourcePath("conf")
	if !file_system.Exists(dirPath) {
		t.Error("not expected")
	}
}

