package uncdn

import (
	"strings"
	"testing"
)

func TestTrumbowyg2310Css(t *testing.T) {
	output := Trumbowyg2310Css()
	expected := "Trumbowyg"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestTrumbowyg2310Js(t *testing.T) {
	output := Trumbowyg2310Js()
	expected := "Trumbowyg"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
