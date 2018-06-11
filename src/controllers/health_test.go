package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AppHealth)
	handler.ServeHTTP(rr, req)

	expected := `{"status":"OK"}`
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, strings.TrimSuffix(rr.Body.String(), "\n"))
}