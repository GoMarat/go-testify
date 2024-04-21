package main

import (
		"net/http"
	    "net/http/httptest"
		"testing"
		"strings"

		"github.com/stretchr/testify/assert"
		"github.com/stretchr/testify/require"
	)

	func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
		totalCount := 4
		req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)
	
		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)

			body := responseRecorder.Body.String()
		list := strings.Split(body, ",")

			assert.Equal(t, len(list), totalCount)
	}

	func TestMainHandlerWhenOk(t *testing.T) {
		req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req) 

		status := responseRecorder.Code
		body := responseRecorder.Body.String()

		require.Equal(t, status, http.StatusOK)
		assert.NotEmpty(t, body)
	}

	func TestMainHandlerWhenCityIsWrong(t *testing.T) {
		req := httptest.NewRequest("GET", "/cafe?count=4&city=Chelyabinsk", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req) 

		status := responseRecorder.Code
		body := responseRecorder.Body.String()

		require.Equal(t, status, http.StatusBadRequest)
		assert.Equal(t, "wrong city value", body)
	}