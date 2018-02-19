// https://gobyexample.com/reading-files

package file_io

import (
	"testing"
	"io/ioutil"
	"os"
	"io"
	"bufio"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func TestReadFileToByteArray_ExpectFirstByte(t *testing.T) {
	filePath := resourcePath("aliases")
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	if 35 != dat[0] {  // #
		t.Error()
	}
}


func TestReadFile_ExpectString(t *testing.T) {
	filePath := resourcePath("aliases")
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	if "# python\na" != string(dat[:10]) {
		t.Error()
	}
}


func TestOpenFileHandle_ExpectSuccessful(t *testing.T) {
	filePath := resourcePath("aliases")
	_, err := os.Open(filePath)
	check(err)
}


func TestReadFromFileHandle_FixedBufferSize(t *testing.T) {
	f, err := os.Open(resourcePath("aliases"))
	check(err)
	b1 := make([]byte, 1005)
	n1, err := f.Read(b1)
	check(err)
	if 483 != n1 {
		t.Error(n1)
	}
}


func TestReadFromFileHandle_Seek(t *testing.T) {
	f, _ := os.Open(resourcePath("aliases"))
	f.Seek(155, 0)
	b2 := make([]byte, 5)
	_, err := f.Read(b2)
	check(err)
	if " pych" != string(b2) {
		t.Error()
	}
}


func TestReadAtLeast_ExpectActualLength(t *testing.T) {
	f, _ := os.Open(resourcePath("aliases"))
	f.Seek(480, 0)
	b3 := make([]byte, 10)
	length, _ := io.ReadAtLeast(f, b3, 10)
	if 3 != length {
		t.Error()
	}
}


func TestBufIOReader_CloseWithDefer(t *testing.T) {
	f, _ := os.Open(resourcePath("aliases"))
	defer f.Close()
	reader := bufio.NewReader(f)
	buf, err := reader.Peek(5)
	check(err)
	if "# pyt" != string(buf) {
		t.Error()
	}
}


