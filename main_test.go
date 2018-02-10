package main

import (
	"net/http/httptest"
	"testing"
)

func TestNegative(t *testing.T) {
	rr := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/fibonacci/-4", nil)
	router := setupRouter()

	router.ServeHTTP(rr, request)
	if status := rr.Code; status != 400 {
		t.Errorf("Expected 400, instead got %v", status)
	}

}

func TestAlpha(t *testing.T) {
	rr := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/fibonacci/afe", nil)
	router := setupRouter()

	router.ServeHTTP(rr, request)
	if status := rr.Code; status != 400 {
		t.Errorf("Expected 400, instead got %v", status)
	}

}

func TestUnicode(t *testing.T) {
	rr := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/fibonacci/™°ñ", nil)
	router := setupRouter()

	router.ServeHTTP(rr, request)
	if status := rr.Code; status != 400 {
		t.Errorf("Expected 400, instead got %v", status)
	}

}

//TODO: Mock the Fibonacci function to ensure it runs properly
