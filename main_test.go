package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RedHatInsights/consoledot-go-starter-app/providers"
	"github.com/RedHatInsights/consoledot-go-starter-app/routes"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

type mockQueryRow struct {
	name string
}

func (m *mockQueryRow) Scan(dest ...interface{}) error {
	retVal := "Database : starter-app-db, User : shadowman"
	*dest[0].(*string) = retVal
	return nil
}

type mockConnectionPool struct {
	name string
}

func (m *mockConnectionPool) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return &mockQueryRow{}
}
func (m *mockConnectionPool) Close() {
}

var (
	mockProviderManager = providers.Providers{
		DBConnectionPool: &mockConnectionPool{
			name: "mockPool",
		},
	}
)

func TestReadinessProbe(t *testing.T) {
	router := routes.SetupRouter(apiPath, mockProviderManager)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/readyz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"ready\":\"OK\"}", w.Body.String())
}

func TestLivlinessProbeRoute(t *testing.T) {
	router := routes.SetupRouter(apiPath, mockProviderManager)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/livez", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"alive\":\"OK\"}", w.Body.String())
}

func TestHelloRoute(t *testing.T) {
	router := routes.SetupRouter(apiPath, mockProviderManager)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", apiPath+"/v1/hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"hello\":\"world\"}", w.Body.String())
}

func TestDBInfoRoute(t *testing.T) {
	router := routes.SetupRouter(apiPath, mockProviderManager)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", apiPath+"/v1/db-info", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"db-info\":\"Database : starter-app-db, User : shadowman\"}", w.Body.String())
}
