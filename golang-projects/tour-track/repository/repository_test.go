package repository

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchReturnsHTTPStatusError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "unavailable", http.StatusServiceUnavailable)
	}))
	defer server.Close()

	client := &Client{HTTP: server.Client(), BaseURL: server.URL}
	_, _, err := client.Fetch(context.Background())
	if err == nil {
		t.Fatal("Fetch() error = nil, want HTTP status error")
	}
	if got := fmt.Sprint(err); got == "" {
		t.Fatal("Fetch() returned an empty error")
	}
}

func TestCountryFromLocation(t *testing.T) {
	tests := []struct {
		location string
		want     string
	}{
		{location: "paris-france", want: "france"},
		{location: "new_york-us", want: "us"},
		{location: "  berlin - germany  ", want: "germany"},
		{location: "", want: ""},
	}

	for _, tt := range tests {
		if got := CountryFromLocation(tt.location); got != tt.want {
			t.Errorf("CountryFromLocation(%q) = %q, want %q", tt.location, got, tt.want)
		}
	}
}
