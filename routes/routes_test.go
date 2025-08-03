package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRouter_ReturnsRouter(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/notfound", nil)

	router.ServeHTTP(w, req)

	if w.Code == 0 {
		t.Errorf("Expected HTTP code, got 0")
	}
}
