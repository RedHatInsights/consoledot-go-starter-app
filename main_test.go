package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RedHatInsights/consoledot-go-starter-app/routes"
	"github.com/stretchr/testify/assert"
)

func TestReadinessProbe(t *testing.T) {
	router := routes.SetupRouter(apiPath)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"ready\":\"OK\"}", w.Body.String())
}

func TestLivlinessProbeRoute(t *testing.T) {
	router := routes.SetupRouter(apiPath)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/livez", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"alive\":\"OK\"}", w.Body.String())
}

func TestHelloRoute(t *testing.T) {
	router := routes.SetupRouter(apiPath)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", apiPath+"/v1/hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"hello\":\"world\"}", w.Body.String())
}
