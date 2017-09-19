package stringOps

import (
	"strings"
	"strconv"
)


func IsUpper(someStr string) bool {
	if len(someStr) == 1 && someStr == strings.ToUpper(someStr) {
		return true
	}
	return false
}


func IsDigit(someStr string) bool {
	if _, err := strconv.Atoi(someStr); err == nil {
		return true
	}
	return false
}