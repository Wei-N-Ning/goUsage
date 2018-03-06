package fileSystem

import (
	"testing"
	"path"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Fatalf("Expected: %s; Got: %s", expected, actual)
	}
}

func TestGetBaseName(t *testing.T) {
	assertEqual(t, path.Base(""), ".")
	assertEqual(t, path.Base("/doom/e1/m1/imp.id1"), "imp.id1")

	// can not handle windows file path (go 1.8)
	assertEqual(t, path.Base(`C:\program files\doom\e1\m1\imp.id1`), `C:\program files\doom\e1\m1\imp.id1`)
}
