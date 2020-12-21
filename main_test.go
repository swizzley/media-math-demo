package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testServer() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/init", initDB)
	return r
}

func TestInitDB(t *testing.T) {

	endpoint := "/init"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testServer().ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
