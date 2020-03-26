package requests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/serboox/statsd-notifier/src/configs"

	"github.com/stretchr/testify/assert"
)

func TestNoRoute(t *testing.T) {
	ctx := configs.NewContextMock()
	router := SetupRouter(ctx)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/1234", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), w.Body.String())
}
