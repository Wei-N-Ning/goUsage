package projectManagement

import (
	"testing"
)


func TestGetSourceLocation(t *testing.T) {
	filePath := SourceLocation()
	if len(filePath) == 0 {
		t.Error("not expected")
	}
}
