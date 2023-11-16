package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nanpangyou/pixel-backend-go/internal/router"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := router.New()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong1pong2", w.Body.String())
}
