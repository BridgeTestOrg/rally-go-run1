package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomePageRenders(t *testing.T) {
	rec := httptest.NewRecorder()
	NewRouter().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "<h1>Rally</h1>") {
		t.Fatalf("body does not contain the Rally heading")
	}
}
