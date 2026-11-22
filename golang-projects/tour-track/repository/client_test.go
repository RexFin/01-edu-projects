package repository

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFetchLoadsAndJoinsAllAPIResources(t *testing.T) {
	seen := make(map[string]bool)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen[r.URL.Path] = true
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/artists":
			fmt.Fprint(w, `{"index":[{"id":1,"name":"Example","creationDate":2000}]}`)
		case "/locations":
			fmt.Fprint(w, `{"index":[{"id":1,"locations":["paris-france","london-uk"]}]}`)
		case "/dates":
			fmt.Fprint(w, `{"index":[{"id":1,"dates":["01-01-2020","02-01-2020"]}]}`)
		case "/relation":
			fmt.Fprint(w, `{"index":[{"id":1,"datesLocations":{"paris-france":["01-01-2020"],"london-uk":["02-01-2020"]}}]}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := &Client{HTTP: server.Client(), BaseURL: server.URL}
	artists, relations, err := client.Fetch(context.Background())
	if err != nil {
		t.Fatalf("Fetch() error = %v", err)
	}
	if len(artists) != 1 || artists[0].Name != "Example" {
		t.Fatalf("unexpected artists: %#v", artists)
	}
	relation := relations[1]
	if !reflect.DeepEqual(relation.Locations, []string{"paris-france", "london-uk"}) {
		t.Errorf("locations = %#v", relation.Locations)
	}
	if !reflect.DeepEqual(relation.Dates, []string{"01-01-2020", "02-01-2020"}) {
		t.Errorf("dates = %#v", relation.Dates)
	}
	if len(relation.DatesLocations) != 2 {
		t.Errorf("relation map = %#v", relation.DatesLocations)
	}
	for _, path := range []string{"/artists", "/locations", "/dates", "/relation"} {
		if !seen[path] {
			t.Errorf("API resource %s was not requested", path)
		}
	}
}
