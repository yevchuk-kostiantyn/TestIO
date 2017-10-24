package rest

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestcheckEnteredCredentials(t *testing.T) {
	req, err := http.NewRequest("PATCH", "/login", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checkEnteredCredentials)

	handler.ServeHTTP(rr, req)

}

func TestgetNewUser(t *testing.T) {

}

func TestsendEmail(t *testing.T) {

}