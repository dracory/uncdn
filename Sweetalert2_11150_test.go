package uncdn

import (
	"strings"
	"testing"
)

func TestSweetalert2_11150(t *testing.T) {
	output := Sweetalert2_11150()
	expected := "SweetAlert2"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
