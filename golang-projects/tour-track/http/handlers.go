package http

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"tour-track/internal/activity"
	"tour-track/repository"
)

type Handler struct {
	API       *repository.Client
	Templates *template.Template
}
type ArtistView struct {
	Artist    repository.Artist
	Score     int
	Activity  string
	Rising    bool
	Concerts  int
	Countries int
	Relation  repository.Relation
}
type PageData struct {
	Artists  []ArtistView
	Artist   ArtistView
	Error    string
	Notice   string
	YearFrom int
	YearTo   int
}

func New(api *repository.Client, t *template.Template) *Handler {
	return &Handler{API: api, Templates: t}
}
func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/artist", h.Artist)
}

func (h *Handler) load(r *http.Request) ([]ArtistView, bool, error) {
	ctx, cancel := context.WithTimeout(r.Context(), 12*time.Second)
	defer cancel()
	artists, relations, err := h.API.Fetch(ctx)
	fallback := false

	if err != nil {
		log.Printf("Groupie API unavailable, using local preview data: %v", err)
		artists, relations = repository.FallbackData()
		fallback = true
	}
	now := time.Now()
	year := now.Year()
	views := make([]ArtistView, 0, len(artists))
	for _, artist := range artists {
		relation := relations[artist.ID]
		countries := make(map[string]struct{})
		locations := relation.Locations
		concertDates := relation.Dates
		if len(locations) == 0 {
			for location, dates := range relation.DatesLocations {
				locations = append(locations, location)
				concertDates = append(concertDates, dates...)
			}
		}
		if len(concertDates) == 0 {
			for _, dates := range relation.DatesLocations {
				concertDates = append(concertDates, dates...)
			}
		}
		for _, location := range locations {
			country := strings.ToLower(repository.CountryFromLocation(location))
			if country != "" {
				countries[country] = struct{}{}
			}
		}
		recent := 0
		for _, raw := range concertDates {
			if isRecent(raw, now) {
				recent++
			}
		}
		concerts := len(concertDates)
		years := year - artist.CreationDate
		if years < 0 {
			years = 0
		}
		score, label := activity.Classify(len(countries), concerts, years)
		views = append(views, ArtistView{Artist: artist, Score: score, Activity: label, Rising: activity.IsRisingStar(recent, concerts, years), Concerts: concerts, Countries: len(countries), Relation: relation})
	}
	return views, fallback, nil
}
func isRecent(raw string, now time.Time) bool {
	// The API prefixes some entries with '*' to mark dates; the marker is not part of the date.
	raw = strings.TrimPrefix(raw, "*")
	for _, layout := range []string{"02-01-2006", "2006-01-02", "2006-1-2"} {
		if d, err := time.Parse(layout, raw); err == nil {
			return !d.Before(now.AddDate(-2, 0, 0)) && !d.After(now)
		}
	}
	return false
}
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.renderError(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		h.renderError(w, http.StatusMethodNotAllowed)
		return
	}
	views, fallback, err := h.load(r)
	if err != nil {
		log.Printf("load artist data for home page: %v", err)
		h.renderError(w, http.StatusBadGateway)
		return
	}
	from, _ := strconv.Atoi(r.URL.Query().Get("year_from"))
	to, _ := strconv.Atoi(r.URL.Query().Get("year_to"))
	filtered := views[:0]
	for _, v := range views {
		if from > 0 && v.Artist.CreationDate < from {
			continue
		}
		if to > 0 && v.Artist.CreationDate > to {
			continue
		}
		filtered = append(filtered, v)
	}
	sort.Slice(filtered, func(i, j int) bool {
		return strings.ToLower(filtered[i].Artist.Name) < strings.ToLower(filtered[j].Artist.Name)
	})
	notice := ""
	if fallback {
		notice = "The live artist API is unavailable. Showing the built-in preview dataset."
	}
	h.render(w, "index.html", PageData{Artists: filtered, YearFrom: from, YearTo: to, Notice: notice})
}
func (h *Handler) Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.renderError(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		h.renderError(w, http.StatusBadRequest)
		return
	}
	views, fallback, err := h.load(r)
	if err != nil {
		log.Printf("load artist data for artist page: %v", err)
		h.renderError(w, http.StatusBadGateway)
		return
	}
	for _, v := range views {
		if v.Artist.ID == id {
			notice := ""
			if fallback {
				notice = "The live artist API is unavailable. Showing the built-in preview dataset."
			}
			h.render(w, "artists.html", PageData{Artist: v, Notice: notice})
			return
		}
	}
	h.renderError(w, http.StatusNotFound)
}
func (h *Handler) render(w http.ResponseWriter, name string, data PageData) {
	var buf bytes.Buffer
	if err := h.Templates.ExecuteTemplate(&buf, name, data); err != nil {
		h.renderError(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = buf.WriteTo(w)
}
func (h *Handler) renderError(w http.ResponseWriter, code int) {
	var buf bytes.Buffer
	data := map[string]any{"Status": code, "Message": http.StatusText(code)}
	if err := h.Templates.ExecuteTemplate(&buf, "error.html", data); err != nil {
		http.Error(w, http.StatusText(code), code)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	_, _ = buf.WriteTo(w)
}
