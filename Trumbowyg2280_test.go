package uncdn

import (
	"strings"
	"testing"
)

func TestTrumbowyg2280Css(t *testing.T) {
	output := Trumbowyg2280Css()
	expected := "Trumbowyg"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestTrumbowyg2280Js(t *testing.T) {
	output := Trumbowyg2280Js()
	expected := "Trumbowyg"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
