package uncdn

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestnewCdnController verifies controller can be created
func TestCdnControllerConstructor(t *testing.T) {
	controller := newCdnController()
	if controller == nil {
		t.Error("newCdnController() returned nil")
	}
}

// TestHandlerEmptyRequest verifies handler with empty request
func TestHandlerEmptyRequest(t *testing.T) {
	controller := newCdnController()

	req := httptest.NewRequest(http.MethodGet, "/cdn/", nil)
	w := httptest.NewRecorder()

	controller.Handler(w, req)

	resp := w.Result()
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body := w.Body.String()
	if body != "Nothing requested" {
		t.Errorf("Expected 'Nothing requested', got '%s'", body)
	}
}

// TestHandlerNoExtension verifies handler with no extension (trailing dot)
func TestHandlerNoExtension(t *testing.T) {
	controller := newCdnController()

	// A name ending with . has empty extension after the dot
	req := httptest.NewRequest(http.MethodGet, "/cdn/jq360.", nil)
	w := httptest.NewRecorder()

	controller.Handler(w, req)

	resp := w.Result()
	defer func() { _ = resp.Body.Close() }()

	body := w.Body.String()
	expected := "No extension provided"
	if body != expected {
		t.Errorf("Expected '%s', got '%s'", expected, body)
	}
}

// TestHandlerUnsupportedExtension verifies handler with unsupported extension
func TestHandlerUnsupportedExtension(t *testing.T) {
	controller := newCdnController()

	req := httptest.NewRequest(http.MethodGet, "/cdn/jq360.html", nil)
	w := httptest.NewRecorder()

	controller.Handler(w, req)

	resp := w.Result()
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body := w.Body.String()
	expected := "Extension html not supported"
	if body != expected {
		t.Errorf("Expected '%s', got '%s'", expected, body)
	}
}

// TestHandlerJSRequest verifies JS compilation endpoint
func TestHandlerJSRequest(t *testing.T) {
	controller := newCdnController()

	req := httptest.NewRequest(http.MethodGet, "/cdn/jq360.js", nil)
	w := httptest.NewRecorder()

	controller.Handler(w, req)

	resp := w.Result()
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// JS requests return proper content type
	if resp.Header.Get("Content-Type") != "application/javascript" {
		t.Errorf("Expected application/javascript content type, got %s", resp.Header.Get("Content-Type"))
	}

	// Verify Content-Encoding is gzip
	if resp.Header.Get("Content-Encoding") != "gzip" {
		t.Errorf("Expected gzip content encoding, got %s", resp.Header.Get("Content-Encoding"))
	}

	// Verify body is not empty
	if w.Body.Len() == 0 {
		t.Error("Expected non-empty response body")
	}
}

// TestHandlerCSSRequest verifies CSS compilation endpoint
func TestHandlerCSSRequest(t *testing.T) {
	controller := newCdnController()

	req := httptest.NewRequest(http.MethodGet, "/cdn/bs523.css", nil)
	w := httptest.NewRecorder()

	controller.Handler(w, req)

	resp := w.Result()
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// CSS requests return proper content type
	if resp.Header.Get("Content-Type") != "text/css" {
		t.Errorf("Expected text/css content type, got %s", resp.Header.Get("Content-Type"))
	}

	// Verify Content-Encoding is gzip
	if resp.Header.Get("Content-Encoding") != "gzip" {
		t.Errorf("Expected gzip content encoding, got %s", resp.Header.Get("Content-Encoding"))
	}

	// Verify body is not empty
	if w.Body.Len() == 0 {
		t.Error("Expected non-empty response body")
	}
}
