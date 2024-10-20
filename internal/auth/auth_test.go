package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	test_header := http.Header{
		"Authorization": []string{"ApiKey your_key"},
	}

	key, err := GetAPIKey(test_header)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedKey := "your_key"
	if key != expectedKey {
		t.Fatalf("expected: %v, got: %v", expectedKey, key)
	}
}
