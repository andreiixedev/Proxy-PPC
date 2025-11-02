package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestBuildsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	RegisterBuildsRoute(r)

	req, _ := http.NewRequest("GET", "/client/builds.json", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}
