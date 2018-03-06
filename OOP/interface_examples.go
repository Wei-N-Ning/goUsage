package OOP

import "strings"

type GenereratorI interface{
	Gen(s string) string
}

type ToUpper struct {
	value int
}

func (this *ToUpper) Gen(s string) string {
	return strings.ToUpper(s)
}

type ToLower struct {
	value int
}

func (this *ToLower) Gen(s string) string {
	return strings.ToLower(s)
}