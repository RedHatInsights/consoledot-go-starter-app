package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RedHatInsights/consoledot-go-starter-app/routes"
	"github.com/stretchr/testify/assert"
)

func TestReadinessProbe(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/probes/ready", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"ready\":\"OK\"}", w.Body.String())
}

func TestLivlinessProbeRoute(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/probes/alive", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"alive\":\"OK\"}", w.Body.String())
}

func TestHelloRoute(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"hello\":\"world\"}", w.Body.String())
}
