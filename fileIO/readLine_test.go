package fileIO

import (
	"testing"
	"os"
	"bufio"
	"github.com/powergun/goUsage/testsWriting"
	"strings"
)


func TestReadLineFromFile(t *testing.T) {
	filePath := resourcePath("aliases")
	f, err := os.Open(filePath)
	testsWriting.Check(err)
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