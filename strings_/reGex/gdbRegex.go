package reGex

import (
	"strings"
	"errors"
	"regexp"
)

type parser struct {
	line         string
	tokens       []string
	currentIndex int
	lastIndex    int
	// Id, Address, Subroutine, Arguments, BinaryFile, SourceFile
	m            map[string]string
}

func newParser(line string) (*parser, bool) {
	l := strings.TrimLeft(line, " \n\t\r")
	l = strings.TrimRight(l, " \n\t\r")
	if l == "" {
		return nil, false
	}
	tokens := strings.Split(l, " ")
	if len(tokens) == 0{
		return nil, false
	}
	p := parser{
		line: line,
		tokens: tokens,
		currentIndex: 0,
		lastIndex: len(tokens) - 1,
		m: make(map[string]string),
	}
	return &p, true
}

func (this *parser) currentToken() string {
	return strings.TrimSpace(this.tokens[this.currentIndex])
}

func (this *parser) skipToken(token string) error {
	for {
		if this.currentIndex > this.lastIndex {
			return errors.New("incomplete line")
		}
		if this.currentToken() != token {
			break
		}
		this.currentIndex++
	}
	return nil
}

func (this *parser) findId() error {
	for {
		if this.currentIndex > this.lastIndex {
			return errors.New("invalid line: missing #")
		}
		if token := this.currentToken(); strings.HasPrefix(token, "#") {
			this.m["Id"] = strings.TrimPrefix(token, "#")
			break
		}
		this.currentIndex++
	}
	this.currentIndex++
	return nil
}

// three scenarios can happen:
// ?? (...)
// call_function (...)
// address in something (...)
func (this *parser) findAddress() {
	token := this.currentToken()
	addr := ""
	if strings.HasPrefix(token, "0x") {
		addr = token
		this.currentIndex++
	}
	this.m["Address"] = addr
}

func (this *parser) findSubroutine() error {
	err := this.skipToken("in")
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`\?\?|[^ ]+`)
	token := this.currentToken()
	if re.MatchString(token) {
		this.m["Subroutine"] = token
		this.currentIndex++
		return nil
	}
	return errors.New("can not find subroutine")
}

func (this *parser) findArguments() error {
	token := this.currentToken()
	if regexp.MustCompile(`\(.*\)`).MatchString(token) {
		this.m["Arguments"] = token
		this.currentIndex++
		return nil
	}
	args := make([]string, 1)
	if strings.HasPrefix(token, "(") {
		args = append(args, token)
	}
	this.currentIndex++
	for {
		if this.currentIndex > this.lastIndex {
			return errors.New("unclosed argument list")
		}
		token = this.currentToken()
		if token != "" {
			args = append(args, token)
		}
		if strings.HasSuffix(token, ")") {
			break
		}
		this.currentIndex++
	}
	this.m["Arguments"] = strings.Join(args, "")
	this.currentIndex++
	return nil
}

func (this *parser) findBinaryOrSource() {
	this.m["BinaryFile"] = ""
	this.m["SourceFile"] = ""
	for {
		if this.currentIndex > this.lastIndex {
			return
		}
		token := this.currentToken()
		if token == "from" {
			this.currentIndex++
			this.m["BinaryFile"] = this.binaryFilePath()
			return
		}
		if token == "at" {
			this.currentIndex++
			this.m["SourceFile"] = this.sourceFilePath()
			return
		}
		this.currentIndex++
	}
}

func (this *parser) binaryFilePath() string {
	if this.currentIndex > this.lastIndex {
		return ""
	}
	return this.currentToken()
}

func (this *parser) sourceFilePath() string {
	if this.currentIndex > this.lastIndex {
		return ""
	}
	return this.currentToken()
}

func (this *parser) do() error {
	err := this.findId()
	if err != nil {
		return err
	}
	err = this.skipToken("")
	if err != nil {
		return err
	}
	this.findAddress()
	this.findSubroutine()
	err = this.findArguments()
	if err != nil {
		return err
	}
	this.findBinaryOrSource()
	return nil
}

// return: map, bool(isEmptyLine), error
func Parse(line string) (map[string]string, bool, error) {
	parser, OK := newParser(line)
	if !OK {
		return nil, true, nil
	}
	err := parser.do()
	if err != nil {
		return nil, false, err
	}
	return parser.m, false, nil
}
