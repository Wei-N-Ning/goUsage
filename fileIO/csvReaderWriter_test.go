package fileIO

import (
	"encoding/csv"
	"strings"
	"testing"
	"bytes"
)

func TestReadFromCSV(t *testing.T) {
	const text = `
title,year,platform,mem
doom,1993,dos,4mb
`
	str := strings.NewReader(text)
	reader := csv.NewReader(str)
	lines, err := reader.ReadAll()
	assertTrue(t, err == nil)
	assertTrue(t, len(lines) == 2)
	assertStringEqual(t, "platform", lines[0][2])
	assertStringEqual(t, "1993", lines[1][1])
}

func TestWriteToCSV(t *testing.T) {
	header := []string{"title", "year", "platform"}
	productLine := []string{"dune", "1991", "dos"}
	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	writer.Write(header)
	writer.Write(productLine)
	writer.Flush()
	assertStringEqual(t, "title,year,platform\ndune,1991,dos\n", buf.String())
}