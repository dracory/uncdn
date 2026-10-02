package uncdn

import (
	"strings"
	"testing"
)

func TestJquerySmartMenus111(t *testing.T) {
	if !strings.Contains(JquerySmartMenusCss111(), ".sm") {
		t.Error("JquerySmartMenusCss111 failed")
	}
	if !strings.Contains(JquerySmartMenusCssBlueTheme111(), ".sm") {
		t.Error("JquerySmartMenusCssBlueTheme111 failed")
	}
	if !strings.Contains(JquerySmartMenusCssBootstrap4AddOn111(), ".navbar-nav") {
		t.Error("JquerySmartMenusCssBootstrap4AddOn111 failed")
	}
	if !strings.Contains(JquerySmartMenusCssSimpleTheme111(), ".sm") {
		t.Error("JquerySmartMenusCssSimpleTheme111 failed")
	}
	if !strings.Contains(JquerySmartMenusJs111(), "SmartMenus") {
		t.Error("JquerySmartMenusJs111 failed")
	}
}
