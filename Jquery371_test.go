package uncdn

import (
	"strings"
	"testing"
)

func TestJquery371(t *testing.T) {
	output := Jquery371()
	expected := "jQuery v3.7.1"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
