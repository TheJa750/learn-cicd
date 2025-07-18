package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	apiKey := "IHaveToPeeSoBadly"
	correctHeader := make(http.Header)
	missingAPIKey := make(http.Header)
	noAuthHeader := make(http.Header)
	blankAPIKey := make(http.Header)

	correctHeader.Set("Authorization", "ApiKey "+apiKey)
	missingAPIKey.Set("Authorization", apiKey)
	noAuthHeader.Set("Content-Type", "text/plain")
	blankAPIKey.Set("Authorization", "ApiKey ")

	tests := map[string]struct {
		input        http.Header
		want         string
		expected_err error
	}{
		"correct header":   {input: correctHeader, want: apiKey, expected_err: nil},
		"missing 'ApiKey'": {input: missingAPIKey, want: "", expected_err: ErrNoAuthHeaderIncluded},
		"wrong header Key": {input: noAuthHeader, want: "", expected_err: ErrNoAuthHeaderIncluded},
		"blank API Key":    {input: blankAPIKey, want: "", expected_err: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if got != tc.want || err != tc.expected_err {
				t.Fatalf("expected string: %v, got string: %v, expected error: %v, got error: %v", tc.want, got, tc.expected_err, err)
			}
		})
	}
}
