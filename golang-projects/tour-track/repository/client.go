package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const apiBaseURL = "https://groupietrackers.01edu.ai/api"

// Artist is the data returned by the /artists endpoint.
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// Relation links locations and dates to an artist.
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
	Locations      []string            `json:"-"`
	Dates          []string            `json:"-"`
}

type artistsResponse struct {
	Index []Artist `json:"index"`
}

func (r *artistsResponse) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return fmt.Errorf("artists response is empty")
	}

	switch data[0] {
	case '[':
		return json.Unmarshal(data, &r.Index)
	case '{':
		type wrappedArtistsResponse artistsResponse
		return json.Unmarshal(data, (*wrappedArtistsResponse)(r))
	default:
		return fmt.Errorf("artists response must be an array or an object with an index field")
	}
}

type locationsResponse struct {
	Index []locationData `json:"index"`
}

type locationData struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type datesResponse struct {
	Index []dateData `json:"index"`
}

type dateData struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type relationsResponse struct {
	Index []Relation `json:"index"`
}

// Client loads data from the Groupie Tracker API.
type Client struct {
	HTTP    *http.Client
	BaseURL string
}

func NewClient() *Client {
	return &Client{
		HTTP:    &http.Client{Timeout: 8 * time.Second},
		BaseURL: apiBaseURL,
	}
}

// get downloads and decodes one API endpoint.
func (c *Client) get(ctx context.Context, endpoint string, destination any) error {
	if c == nil || c.HTTP == nil {
		return fmt.Errorf("HTTP client is not initialized")
	}
	if strings.TrimSpace(c.BaseURL) == "" {
		return fmt.Errorf("API base URL is empty")
	}

	url := strings.TrimRight(c.BaseURL, "/") + endpoint
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("create request for %s: %w", endpoint, err)
	}

	response, err := c.HTTP.Do(request)
	if err != nil {
		return fmt.Errorf("download %s: %w", endpoint, err)
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("download %s: API returned %s", endpoint, response.Status)
	}
	if err := json.NewDecoder(response.Body).Decode(destination); err != nil {
		return fmt.Errorf("decode %s: %w", endpoint, err)
	}
	return nil
}

// Fetch gets all four resources and joins locations and dates by artist ID.
func (c *Client) Fetch(ctx context.Context) ([]Artist, map[int]Relation, error) {
	var artistData artistsResponse
	if err := c.get(ctx, "/artists", &artistData); err != nil {
		return nil, nil, fmt.Errorf("fetch artists: %w", err)
	}

	var locationData locationsResponse
	if err := c.get(ctx, "/locations", &locationData); err != nil {
		return nil, nil, fmt.Errorf("fetch locations: %w", err)
	}

	var datesData datesResponse
	if err := c.get(ctx, "/dates", &datesData); err != nil {
		return nil, nil, fmt.Errorf("fetch dates: %w", err)
	}

	var relationData relationsResponse
	if err := c.get(ctx, "/relation", &relationData); err != nil {
		return nil, nil, fmt.Errorf("fetch relations: %w", err)
	}

	// Start with an entry per artist so missing API data still produces empty
	// lists instead of nil values.
	relations := map[int]Relation{}
	for _, artist := range artistData.Index {
		relations[artist.ID] = Relation{
			ID:             artist.ID,
			DatesLocations: map[string][]string{},
			Locations:      []string{},
			Dates:          []string{},
		}
	}

	for _, relation := range relationData.Index {
		artistRelation := relations[relation.ID]
		artistRelation.ID = relation.ID
		artistRelation.DatesLocations = relation.DatesLocations
		if artistRelation.DatesLocations == nil {
			artistRelation.DatesLocations = map[string][]string{}
		}
		relations[relation.ID] = artistRelation
	}

	for _, data := range locationData.Index {
		relation := relations[data.ID]
		relation.ID = data.ID
		relation.Locations = data.Locations
		if relation.Locations == nil {
			relation.Locations = []string{}
		}
		relations[data.ID] = relation
	}

	for _, data := range datesData.Index {
		relation := relations[data.ID]
		relation.ID = data.ID
		relation.Dates = data.Dates
		if relation.Dates == nil {
			relation.Dates = []string{}
		}
		relations[data.ID] = relation
	}

	return artistData.Index, relations, nil
}

// CountryFromLocation returns the country from a value such as "london-uk".
func CountryFromLocation(location string) string {
	location = strings.TrimSpace(location)
	separator := strings.LastIndex(location, "-")
	if separator == -1 {
		return location
	}
	return strings.TrimSpace(location[separator+1:])
}
