package fileIO

import "testing"
import (
	"github.com/powergun/goUsage/testsWriting"
	"io/ioutil"
	"os"
	"bufio"
)

var (
	testFile = "/tmp/goUsage_fileIO_writeFile.dat"
)


func TestWriteFileExpectSuccess(t *testing.T) {
	d1 := []byte("something")
	err := ioutil.WriteFile(testFile, d1, 0644)
	testsWriting.Check(err)
	d2, err := ioutil.ReadFile(testFile)
	testsWriting.Check(err)
	if string(d1) != string(d2) {
		t.Error()
	}
}


func TestOpenFileForWriting(t *testing.T) {
	f, err := os.Create(testFile)
	testsWriting.Check(err)
	defer f.Close()
	d1 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n1, err := f.Write(d1)
	testsWriting.Check(err)
	if 9 != n1 {
		t.Error()
	}
	n2, err := f.WriteString("something")
	testsWriting.Check(err)
	if 9 != n2 {
		t.Error()
	}
	f.Sync()
}


func TestOpenBufferObjectForWriting(t *testing.T) {
	f, err := os.Create(testFile)
	testsWriting.Check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	d1 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n1, err := w.Write(d1)
	testsWriting.Check(err)
	if 9 != n1 {
		t.Error()
	}
	n2, err := w.WriteString("something")
	testsWriting.Check(err)
	if 9 != n2 {
		t.Error()
	}
	w.Flush()
}


