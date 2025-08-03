package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthMiddleware_MissingToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("GET", "/", nil)
	c.Request = req

	mw := AuthMiddlewares()
	mw(c)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 Unauthorized, got %d", w.Code)
	}
}
