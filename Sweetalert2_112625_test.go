package uncdn

import (
	"strings"
	"testing"
)

func TestSweetalert2_112625(t *testing.T) {
	output := Sweetalert2_112625()
	expected := "SweetAlert2"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
