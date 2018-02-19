package file_io

import (
	"testing"
	"github.com/powergun/goUsage/testing_"
	"os"
	"bufio"
	"strings"
)


func TestReadLineFromFile(t *testing.T) {
	filePath := resourcePath("aliases")
	f, err := os.Open(filePath)
	testing_.Check(err)
	scanner := bufio.NewScanner(f)
	numAlias := 0
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			numAlias += 1
		}
	}
	if 5 != numAlias {
		t.Error()
	}
}
