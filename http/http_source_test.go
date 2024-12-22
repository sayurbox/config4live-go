package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sayurbox/config4live-go/internal"
	"github.com/stretchr/testify/assert"
)

// Test HandleResponse with Successful Response
func TestHandleResponse_Success(t *testing.T) {
	responseBody := `{
		"success": true,
		"data": {
			"id": "config1",
			"name": "sample-config",
			"value": "12345"
		},
		"error": null
	}`

	recorder := httptest.NewRecorder()
	recorder.WriteHeader(http.StatusOK)
	recorder.WriteString(responseBody)

	resp := recorder.Result()
	defer resp.Body.Close()

	config, err := handleResponse(resp)
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "config1", config.ID)
	assert.Equal(t, "sample-config", config.Name)
}

// Test HandleResponse with Error Response
func TestHandleResponse_Error(t *testing.T) {
	responseBody := `{
		"success": false,
		"data": null,
		"error": "Invalid configuration key"
	}`

	recorder := httptest.NewRecorder()
	recorder.WriteHeader(http.StatusOK)
	recorder.WriteString(responseBody)

	resp := recorder.Result()
	defer resp.Body.Close()

	config, err := handleResponse(resp)
	assert.Nil(t, config)
	assert.Error(t, err)
	assert.Equal(t, "Invalid configuration key", err.Error())
}

// Test HandleResponse with HTTP Error
func TestHandleResponse_HttpError(t *testing.T) {
	responseBody := `{
		"success": true,
		"data": null,
		"error": null
	}`

	recorder := httptest.NewRecorder()
	recorder.WriteHeader(http.StatusInternalServerError)
	recorder.WriteString(responseBody)

	resp := recorder.Result()
	defer resp.Body.Close()

	config, err := handleResponse(resp)
	assert.Nil(t, config)
	assert.Error(t, err)
	assert.Equal(t, "failed response from live-config", err.Error())
}

// Test HttpSource Get Config with Hystrix
func TestHttpSource_Get_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"success": true,
			"data": {
				"id": "config1",
				"name": "test-key",
				"value": "value123"
			}
		}`))
	}))
	defer server.Close()

	source := NewHttpSource(func(s *HttpSource) {
		s.url = server.URL
		s.hystrixParam = &internal.HystrixParams{
			Name:    "test-command",
			Timeout: 1000,
		}
	})

	config, err := source.Get("test-key")
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "test-key", config.Name)
}

// Test HttpSource Get with Not Found Config
func TestHttpSource_Get_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"success": true,
			"data": null
		}`))
	}))
	defer server.Close()

	source := NewHttpSource(func(s *HttpSource) {
		s.url = server.URL
		s.hystrixParam = &internal.HystrixParams{
			Name:    "test-command",
			Timeout: 1000,
		}
	})

	config, err := source.Get("missing-key")
	assert.Nil(t, config)
	assert.NoError(t, err)
}

// Test HttpSource Get with Circuit Breaker Fallback
func TestHttpSource_Get_Fallback(t *testing.T) {
	source := NewHttpSource(func(s *HttpSource) {
		s.url = "http://invalid-url"
		s.hystrixParam = &internal.HystrixParams{
			Name:    "test-command",
			Timeout: 1000,
		}
	})

	config, err := source.Get("any-key")
	assert.Nil(t, config)
	assert.Error(t, err)
}
