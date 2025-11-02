package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID            string `josn:"id`
	OrignalURL    string `json:"original_url`
	ShortURL      string `json:"short_url`
	CreateionDate time.Time `json:"creation_date`
}

var urlDB = make(map[string] URL)

func generateShortURL(OrignalURL string ) string {
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
		ID: id,
		OrignalURL: originalURL,
		ShortURL: shortURL,
		CreateionDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error){
	url, ok := urlDB[id]
	if !ok {
		return URL{},
		errors.New("url not found")
	}
	return url,nil
}

func main() {
	fmt.Println("server is running on port 3000...")
	err := http.ListenAndServe(":3000",nil)
	if err != nil {
		fmt.Println("error in running server", err)
	}
}