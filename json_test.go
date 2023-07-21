package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	recorder := httptest.NewRecorder()
	expectedStatusCode := http.StatusInternalServerError
	expectedMessage := "Internal Server Error"

	respondWithError(recorder, expectedStatusCode, expectedMessage)

	if recorder.Code != expectedStatusCode {
		t.Errorf("Expected status code: %d, but got: %d", expectedStatusCode, recorder.Code)
	}

	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response body: %v", err)
	}

	if response["error"] != expectedMessage {
		t.Errorf("Expected error message: %s, but got: %s", expectedMessage, response["error"])
	}
}

func TestRespondWithJSON(t *testing.T) {
	recorder := httptest.NewRecorder()
	expectedStatusCode := http.StatusOK
	expectedPayload := map[string]string{
		"message": "Success",
	}

	respondWithJSON(recorder, expectedStatusCode, expectedPayload)

	if recorder.Code != expectedStatusCode {
		t.Errorf("Expected status code: %d, but got: %d", expectedStatusCode, recorder.Code)
	}

	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response body: %v", err)
	}

	if !areEqualMaps(expectedPayload, response) {
		t.Errorf("Expected payload: %v, but got: %v", expectedPayload, response)
	}
}

func areEqualMaps(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, val := range m1 {
		if m2Val, found := m2[key]; !found || m2Val != val {
			return false
		}
	}

	return true
}
