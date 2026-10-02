package uncdn

import (
	"strings"
	"testing"
)

func TestJqueryUiJs1133(t *testing.T) {
	output := JqueryUiJs1133()
	expected := "jQuery UI"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqueryUiCss1133(t *testing.T) {
	output := JqueryUiCss1133()
	expected := "jQuery UI"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
