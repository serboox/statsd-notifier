package requests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/serboox/statsd-notifier/src/configs"

	"github.com/stretchr/testify/assert"
)

func TestRootPostRoute(t *testing.T) {
	ctx := configs.NewContextMock()
	router := SetupRouter(ctx)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, http.StatusText(http.StatusOK), w.Body.String())
}
