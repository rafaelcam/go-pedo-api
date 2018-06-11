package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestWhenEndValueIsInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/pedo?start=1&end=n", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPeDo)
	handler.ServeHTTP(rr, req)

	expected := `Invalid end value. Try again.`
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, expected, strings.TrimSuffix(rr.Body.String(), "\n"))
}

func TestWhenStartValueIsInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/pedo?start=a&end=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPeDo)
	handler.ServeHTTP(rr, req)

	expected := `Invalid start value. Try again.`
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, expected, strings.TrimSuffix(rr.Body.String(), "\n"))
}

func TestWhenStartIsOneAndEndIsFifteen(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/pedo?start=1&end=15", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPeDo)
	handler.ServeHTTP(rr, req)

	expected := `["1","2","Pé","4","Do","Pé","7","8","Pé","Do","11","Pé","13","14","PéDo"]`
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, strings.TrimSuffix(rr.Body.String(), "\n"))
}