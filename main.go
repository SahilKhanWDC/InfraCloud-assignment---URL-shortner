package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"time"
)

type URL struct {
	ID            string    `josn:"id`
	OrignalURL    string    `json:"original_url`
	ShortURL      string    `json:"short_url`
	CreateionDate time.Time `json:"creation_date`
}

var urlDB = make(map[string]URL)
var domainCount = make(map[string]int)

func generateShortURL(OrignalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OrignalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	fmt.Println("Encode to string", hash)
	fmt.Println("final hashed string", hash[:8])
	return hash[:8]
}

func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL
	urlDB[id] = URL{
		ID:            id,
		OrignalURL:    originalURL,
		ShortURL:      shortURL,
		CreateionDate: time.Now(),
	}

	parsed, err := url.Parse(originalURL)
	if err == nil {
		domain := parsed.Hostname()
		domainCount[domain]++
	}

	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{},
			errors.New("url not found")
	}
	return url, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "don't worry !! home route is serving perfectly :) ")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	short_URL := createURL(data.URL)
	response := struct {
		ShortURL string `json: "short_url"`
	}{ShortURL: short_URL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "invalid request", http.StatusNotFound)
	}
	http.Redirect(w, r, url.OrignalURL, http.StatusFound)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	type domainMetric struct {
		Domain string `json:"domain"`
		Count  int    `json:"count"`
	}

	var metrics []domainMetric
	for domain, count := range domainCount {
		metrics = append(metrics, domainMetric{Domain: domain, Count: count})
	}

	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].Count > metrics[j].Count
	})

	if len(metrics) > 3 {
		metrics = metrics[:3]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}


func main() {

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)
	http.HandleFunc("/metrics", metricsHandler)

	fmt.Println("server is running on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("error in running server", err)
	}
}
