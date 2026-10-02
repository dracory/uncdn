package uncdn

import (
	"strings"
	"testing"
)

func TestPicoCss206(t *testing.T) {
	output := PicoCss206()
	expected := "Pico CSS"
	if !strings.Contains(output, expected) && !strings.Contains(output, "data-theme") {
		t.Error("Does not contain expected substring, Output:" + output)
	}
}
