package OOP

import "testing"

func handle(g GenereratorI, s string) string {
	return g.Gen(s)
}

func TestExpectUpper(t *testing.T) {
	u := &ToUpper{}
	s := handle(u, "asd")
	if "ASD" != s {
		t.Fatal()
	}
}

func TestExpectLower(t *testing.T) {
	u := &ToLower{}
	s := handle(u, "ASd")
	if "asd" != s {
		t.Fatal()
	}
}
