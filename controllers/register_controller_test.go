package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/anggakrnwn/user-auth-api/config"
	"github.com/anggakrnwn/user-auth-api/database"
	"github.com/anggakrnwn/user-auth-api/models"
	"github.com/anggakrnwn/user-auth-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test") // pastikan ENV=test
	config.LoadEnv()
	database.InitDB()

	// Optional: clear table user
	database.DB.Exec("DELETE FROM users")

	code := m.Run()

	// Optional: cleanup
	database.DB.Exec("DELETE FROM users")
	os.Exit(code)
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return routes.SetupRouter()
}

func TestRegisterSuccess(t *testing.T) {
	router := setupTestRouter()

	// Clean database first
	database.DB.Where("1 = 1").Delete(&models.User{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(`{
		"name": "Test User",
		"username": "testuser",
		"email": "test@example.com",
		"password": "password123"
	}`))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User created successfully")
}
