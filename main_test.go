package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"unicode"
)

const (
	inputFile = "./testdata/test1.md"
	resultFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

func RemoveSpace(input []byte) []byte {
	s := string(input)
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return []byte(string(rr))
}

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result, err := parseContent(input, "")
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	// Strip all white spaces from the two to do a tab and format agnostics comparison

	if !bytes.Equal(RemoveSpace(expected), RemoveSpace(result)) {
		t.Logf("golden: \n%s\n", expected)
		t.Logf("result: \n%s\n", result)
		t.Logf("stripped result: \n%s\n", RemoveSpace(result))
		t.Logf("stripped expected: \n%s\n", RemoveSpace(expected))
		t.Error("Result content does not match golden file")
	}
}

func TestRun(t *testing.T) {
	var mockStdout bytes.Buffer
	if err := run(inputFile, "", &mockStdout, true); err != nil {
		t.Fatal(err)
	}
	resultFile := strings.TrimSpace(mockStdout.String())
	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(RemoveSpace(expected), RemoveSpace(result)) {
		t.Logf("golden: \n%s\n", expected)
		t.Logf("result: \n%s\n", result)
		t.Error("Result content does not match golden file")
	}
	os.Remove(resultFile)
}
