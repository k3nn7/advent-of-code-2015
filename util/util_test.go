package util

import (
	"strings"
	"testing"
)

func TestReadInputForNotExistingFile(t *testing.T) {
	_, err := ReadInput("./not-existing")

	if err == nil {
		t.Errorf("Error should be different than nil")
	}

	if !strings.Contains(err.Error(), "no such file") {
		t.Errorf(`Error message "%s" should contain "no such file" string`, err.Error())
	}
}

func TestReadInputForExistingFile(t *testing.T) {
	content, err := ReadInput("./input.txt")

	if err != nil {
		t.Errorf(`Error should be nil got: "%v"`, err)
	}

	expectedContent := "input content\n"
	if content != expectedContent {
		t.Errorf(`Expecting content: "%s" but got: "%s"`, expectedContent, content)
	}
}
