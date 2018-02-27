package terminal_

import (
	"testing"
	"fmt"
)

// http://grokbase.com/t/gg/golang-nuts/1321s2z98y/go-nuts-printf-in-color
func TestPrintlnInColor(t *testing.T) {
	text := "important"
	a := "\x1b[31;1m" + text + "\x1b[0m"
	augmentedText := fmt.Sprintf(`header something... %s footer`, a)
	fmt.Printf("%q", augmentedText)
}
