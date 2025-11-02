package tests

import (
	"net/http"
	"net/http/httptest"
	"newproxy/routes"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStatusHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	routes.RegisterStatusRoute(r)

	req, _ := http.NewRequest("GET", "/status", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", w.Code)
	}
}
