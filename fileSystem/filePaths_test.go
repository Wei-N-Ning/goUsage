package fileSystem

import (
	"testing"
	"os"
	"io/ioutil"
)


func TestGetCurrentDirectory_ExpectNonEmpty(t *testing.T) {
	curr := os.Args[0]
	if (len(curr) == 0) {
		t.Error("not expected")
	}
}


func TestCheckFileSystemEntityExists(t *testing.T) {
	if !Exists("/tmp") {
		t.Error("not expected")
	}
	if Exists("/cygdrive/c") {
		t.Error("not expected")
	}
}


func TestGetContentsFromDirectory(t *testing.T) {
	fInfoA, _ := ioutil.ReadDir("/tmp")
	//for _, fInfo := range fInfoA {
	//}
	if len(fInfoA) == 0 {
		t.Error("not expected")
	}
}