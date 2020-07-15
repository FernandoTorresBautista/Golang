package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8088/fernando@golang.org",
		nil,
	)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200: got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "gopher fernando") {
		t.Errorf("unexpected body in response: %q", rec.Body.String())
	}
	//t.Errorf("something really wrong")
}

/**
> go test test_tools/main_test.go

*/
