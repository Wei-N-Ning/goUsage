package environments

import (
	"os"
	"strings"
)

func HasEnvVar(someKey string) bool {
	for _, pair := range os.Environ() {
		nameAndValue := strings.Split(pair, "=")
		if someKey == nameAndValue[0] {
			return true
		}
	}
	return false
}