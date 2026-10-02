package uncdn

import (
	"strings"
	"testing"
)

func TestJqueryUiJs1142(t *testing.T) {
	output := JqueryUiJs1142()
	expected := "jQuery UI"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqueryUiCss1142(t *testing.T) {
	output := JqueryUiCss1142()
	expected := "jQuery UI"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
