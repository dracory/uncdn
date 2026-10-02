package uncdn

import (
	"strings"
	"testing"
)

func TestJquery400(t *testing.T) {
	output := Jquery400()
	expected := "jQuery v4.0.0"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
