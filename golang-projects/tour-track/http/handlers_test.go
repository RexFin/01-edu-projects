package http

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"tour-track/repository"
)

func TestIsRecentParsesStarredAPIValues(t *testing.T) {
	now := time.Date(2026, time.November, 2, 12, 0, 0, 0, time.UTC)
	if !isRecent("*03-11-2025", now) {
		t.Fatal("expected starred API date in the last two years to count")
	}
	if isRecent("*01-01-2020", now) {
		t.Fatal("expected old date to be outside the two-year window")
	}
}

func TestHomeRendersArtistsFromAPI(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/artists":
			_, _ = w.Write([]byte(`[{"id":1,"name":"API Artist","creationDate":2000,"image":"https://example.com/artist.jpg"}]`))
		case "/locations", "/dates", "/relation":
			_, _ = w.Write([]byte(`{"index":[]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer upstream.Close()
	templates, err := template.ParseGlob("../templates/*.html")
	if err != nil {
		t.Fatalf("parse templates: %v", err)
	}
	handler := New(&repository.Client{HTTP: upstream.Client(), BaseURL: upstream.URL}, templates)
	recorder := httptest.NewRecorder()

	handler.Home(recorder, httptest.NewRequest(http.MethodGet, "/", nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body = %s", recorder.Code, recorder.Body.String())
	}
	if !strings.Contains(recorder.Body.String(), `data-id="1"`) ||
		!strings.Contains(recorder.Body.String(), `data-name="API Artist"`) {
		t.Fatalf("response does not render API artist: %s", recorder.Body.String())
	}
}

func TestHomeShowsPreviewIfAPIIsUnavailable(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.NotFound(w, r) }))
	defer upstream.Close()
	templates, err := template.ParseGlob("../templates/*.html")
	if err != nil {
		t.Fatalf("parse templates: %v", err)
	}
	handler := New(&repository.Client{HTTP: upstream.Client(), BaseURL: upstream.URL}, templates)
	recorder := httptest.NewRecorder()
	handler.Home(recorder, httptest.NewRequest(http.MethodGet, "/", nil))
	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", recorder.Code)
	}
	body := recorder.Body.String()
	if !strings.Contains(body, "built-in preview dataset") || !strings.Contains(body, "Queen") || !strings.Contains(body, "Gorillaz") || !strings.Contains(body, "Legendary") {
		t.Fatalf("home page is missing preview content: %s", body)
	}
}

func TestFilterRequestReturnsMatchingPreviewArtists(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.NotFound(w, r) }))
	defer upstream.Close()
	templates, err := template.ParseGlob("../templates/*.html")
	if err != nil {
		t.Fatalf("parse templates: %v", err)
	}
	handler := New(&repository.Client{HTTP: upstream.Client(), BaseURL: upstream.URL}, templates)
	recorder := httptest.NewRecorder()
	handler.Home(recorder, httptest.NewRequest(http.MethodGet, "/?year_from=2000", nil))
	body := recorder.Body.String()
	if recorder.Code != http.StatusOK || !strings.Contains(body, "Travis Scott") || strings.Contains(body, "Queen") {
		t.Fatalf("filter response status=%d, body=%s", recorder.Code, body)
	}
}

func TestPreviewProfilesContainAuditorExamples(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.NotFound(w, r) }))
	defer upstream.Close()
	templates, err := template.ParseGlob("../templates/*.html")
	if err != nil {
		t.Fatalf("parse templates: %v", err)
	}
	handler := New(&repository.Client{HTTP: upstream.Client(), BaseURL: upstream.URL}, templates)
	cases := []struct {
		id   string
		want []string
	}{
		{"39", []string{"Gorillaz", "26-03-2001"}},
		{"51", []string{"Foo Fighters", "Dave Grohl", "Rami Jaffee"}},
		{"30", []string{"Travis Scott", "Legendary", "santiago-chile", "turku-finland"}},
	}
	for _, tc := range cases {
		recorder := httptest.NewRecorder()
		handler.Artist(recorder, httptest.NewRequest(http.MethodGet, "/artist?id="+tc.id, nil))
		body := recorder.Body.String()
		if recorder.Code != http.StatusOK {
			t.Errorf("artist %s status = %d, want 200", tc.id, recorder.Code)
			continue
		}
		for _, expected := range tc.want {
			if !strings.Contains(body, expected) {
				t.Errorf("artist %s page is missing %q", tc.id, expected)
			}
		}
	}
}
