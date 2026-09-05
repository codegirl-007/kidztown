package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthz(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d", rec.Code)
	}
	if body := rec.Body.String(); body != "ok\n" {
		t.Fatalf("body: got %q", body)
	}
}

func TestIndex(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d", rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "<canvas") {
		t.Fatalf("missing canvas: %s", body)
	}
	if !strings.Contains(body, "/static/app.js") {
		t.Fatalf("missing static script: %s", body)
	}
}

func TestStatic(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/static/app.js", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d", rec.Code)
	}
	body, err := io.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(body) == 0 {
		t.Fatal("empty static file")
	}
}
