package json_

import (
	"bytes"
	"encoding/json"
	"testing"
)

type SUT struct {
	Year  string `json:"year"`
	Title string `json:"title"`
	IDDQD string `json:"iddqd"`
}

func decode(t *testing.T, s string) SUT {
	decoder := json.NewDecoder(bytes.NewBufferString(s))
	suts := make([]SUT, 1)
	if err := decoder.Decode(&suts); err != nil {
		return SUT{}
	}
	return suts[0]
}

func assertFields(t *testing.T, s SUT, expected map[string]string) {
	if s.Year == expected["Year"] && s.Title == expected["Title"] {
	} else {
		t.Fatal()
	}
}

//////////////////////////////////////////////////////////////////////

func TestGivenEmptyJSONExpectNullStruct(t *testing.T) {
	sut := decode(t, "")
	assertFields(t, sut, map[string]string{"Year": "", "Title": ""})
}

func TestGivenInvalidJSONExpectNullStruct(t *testing.T) {
	const invalidRawJSON = `[{"asd": asdasdsd]`
	sut := decode(t, invalidRawJSON)
	assertFields(t, sut, map[string]string{"Year": "", "Title": ""})
}

func TestGivenValidJSONExpectStruct(t *testing.T) {
	const rawJSON = `[{
		"year": "1993",
		"title": "doom"
	}]`
	sut := decode(t, rawJSON)
	assertFields(t, sut, map[string]string{"Year": "1993", "Title": "doom"})
}

func TestExpectExtraFieldsIgnored(t *testing.T) {
	const rawJSON = `[{
		"year": "1993",
		"title": "doom"
	}]`
	sut := decode(t, rawJSON)
	assertFields(t, sut, map[string]string{"Year": "1993", "Title": "doom", "IDDQD": ""})
}
