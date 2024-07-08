package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type URL struct {
	URL       string    `json:"original_url"`
	CreatedAT time.Time `json:"created_at"`
}

var UrlDB = make(map[string]URL)

func GenerateShortUrl(original string) string {
	hasher := md5.New()
	hasher.Write([]byte(original))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
}

func CreateUrl(original string) string {
	shortUrl := GenerateShortUrl(original)
	UrlDB[shortUrl] = URL{
		URL:       original,
		CreatedAT: time.Now(),
	}
	return fmt.Sprintf("http://localhost:8080/%v", shortUrl)
}

func GetURL(id string) (URL, error) {
	if url, ok := UrlDB[id]; !ok {
		return URL{}, errors.New("the short url for this id doesn't exists")
	} else {
		return url, nil
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Write([]byte("Hello there, welcome to the shortener api"))
		} else {
			w.Write([]byte("Sorry this route doesn't exists available routes: \n / \n POST /short\n /url/{id}"))
		}
	})

	mux.HandleFunc("/url/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if url, err := GetURL(id); err != nil {
			http.Error(w, "Sorry the short url with this id doesn't exists", http.StatusNotFound)
		} else {
			http.Redirect(w, r, url.URL, http.StatusFound)
		}
	})

	mux.HandleFunc("POST /short", func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			URL string `json:"url"`
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Please pass a valid body", http.StatusBadRequest)
			return
		}

		shortUrl := CreateUrl(data.URL)
		response := struct {
			ShortURL string `json:"shortURL"`
		}{ShortURL: shortUrl}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&response)
	})

	s := http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	fmt.Println("Server started on port :8080")
	log.Fatal(s.ListenAndServe())
}
