package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr error
	}{
		{
			name:        "Valid Authorization Header",
			headers:     http.Header{"Authorization": []string{"ApiKey myapikey"}},
			expectedKey: "myapikey",
			expectedErr: nil,
		},
		{
			name:        "No Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:        "Malformed Authorization Header",
			headers:     http.Header{"Authorization": []string{"InvalidHeader"}},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		{
			name:        "Invalid Authorization Header",
			headers:     http.Header{"Authorization": []string{}},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			if key != tc.expectedKey {
				t.Errorf("Expected key: %s, but got: %s", tc.expectedKey, key)
			}

			if err != nil && tc.expectedErr != nil {
				if err.Error() != tc.expectedErr.Error() {
					t.Errorf("Expected error: %v, but got: %v", tc.expectedErr, err)
				}
			} else if err != nil || tc.expectedErr != nil {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedErr, err)
			}
		})
	}
}
