package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerValidRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	responseBody := responseRecorder.Body.String()
	assert.NotEmpty(t, responseBody)
	assert.Len(t, strings.Split(responseBody, ","), 2)
}

func TestMainHandlerWrongCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=name&count=4", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	responseBody := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", responseBody)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=4", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	expectedResponse := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	responseBody := responseRecorder.Body.String()

	assert.Equal(t, expectedResponse, responseBody)
	assert.Len(t, strings.Split(responseBody, ","), totalCount)
}
