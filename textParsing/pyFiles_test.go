package textParsing

import (
	"strings"
	"testing"
)

const pySrcLines = `
test_dofoo
not a foo
doofoo
`

const pySrcLinesWithClass = `
class TestSetNthBit(unittest.TestCase):
    def test_setBit(self):
		...
    def test_unsetBit(self):
		...
`

type PyTestFunc struct {
	name string
	doc  string
}

type PyTestClass struct {
	name  string
	doc   string
	funcs []PyTestFunc
}

func CreatePyTestClass() *PyTestClass {
	funcs := make([]PyTestFunc, 0)
	return &PyTestClass{
		"",
		"",
		funcs,
	}
}

func (this *PyTestClass) IsValid() bool {
	return len(this.name) > 0 && len(this.funcs) > 0
}

func Parse(src string) []*PyTestClass {
	classes := make([]*PyTestClass, 0)
	currentClass := CreatePyTestClass()
	for _, line := range strings.Split(src, "\n") {

	}
	if currentClass.IsValid() {
		classes = append(classes, currentClass)
	}
	return classes
}

func TestExpectFreeTestFunctionIdentified(t *testing.T) {
	entities := Parse(pySrcLines)
	if len(entities) != 1 {
		t.Fatal()
	}
}

func TestExpectTestClassIdentified(t *testing.T) {

}
