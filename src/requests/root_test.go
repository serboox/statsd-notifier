package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/serboox/statsd-notifier/src/configs"

	"github.com/stretchr/testify/assert"
)

func TestRootPostRoute(t *testing.T) {
	ctx := configs.NewContextMock()
	router := SetupRouter(ctx)

	requestBody := new(GeoJSON)
	requestBody.Geo.CityName = "New York City"
	requestBody.Geo.ContinentCode = "NA"
	requestBody.Geo.CountryIsoCode = "US"

	reqBody, _ := json.Marshal(requestBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, http.StatusText(http.StatusOK), w.Body.String())
}
