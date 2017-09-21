package environ

import (
	"testing"
	"os"
)


func TestGetEnvironmentVariables(t *testing.T) {
	if os.Getenv("NONEXISTINGFOOBAR") != "" {
		t.Error("not expected")
	}
	if os.Getenv("HOME") == "" {
		t.Error("not expected")
	}
}


func TestSetEnvironmentVariables(t *testing.T) {
	if err := os.Setenv("NONEXISTINGFOOBAR", "foobar"); err != nil {
		t.Error("not expected")
	}
	if err := os.Setenv("", "foobar"); err == nil {
		t.Error("not expected")
	}
	if v := os.Getenv("NONEXISTINGFOOBAR"); v != "foobar" {
		t.Error("not expected")
	}
}


func TestEnvironmentVariableExists(t *testing.T) {
	if !HasEnvVar("HOME") {
		t.Error("not expected")
	}
	if HasEnvVar(">>>>>") {
		t.Error("not expected")
	}
}