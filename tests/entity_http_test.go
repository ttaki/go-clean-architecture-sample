package tests

import (
	"bytes"
	"encoding/json"
	sw "go-clean-architecture/api/openapi"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-clean-architecture/tests"
)

func TestPingEntityRoute(t *testing.T) {
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/entities", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "null", w.Body.String())
}

func TestCreateEntityInvalidRoute(t *testing.T) {
	router := sw.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/entities", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	// assert.Equal(t, "null", w.Body.String())

	requestData, _ := json.Marshal(tests.InvalidEntity{})
	req, _ = http.NewRequest("POST", "/entities", bytes.NewBuffer(requestData))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestCreateEntityRoute(t *testing.T) {
	router := sw.NewRouter()

	w := httptest.NewRecorder()

	requestData, _ := json.Marshal(tests.Entity{Text: "test text"})

	req, _ := http.NewRequest("POST", "/entities", bytes.NewBuffer(requestData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	//  for bind json
	responseData := Entity{}

	responseBytes := []byte(w.Body.String())
	// a, _ := w.Body.ReadBytes(10)
	json.Unmarshal(responseBytes, &responseData)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "test text", responseData.Text)
}
