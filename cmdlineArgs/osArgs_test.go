// https://gobyexample.com/command-line-arguments

package cmdlineArgs

import (
	"testing"
	"os"
)

func TestAccessCmdlineArgs(t *testing.T) {
	if len(os.Args) == 0 {
		t.Fatal()
	}
}
