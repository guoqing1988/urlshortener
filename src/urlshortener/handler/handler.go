package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
	"urlshortener/storage"
)

const ServerHost string = "localhost"
const ServerPort string = "8080"

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method - only POST allowed.", 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body.", 503)
		return
	}

	rawUrl := string(body)
	log.Println("Shorten Handler", rawUrl)

	url, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		http.Error(w, "Invalid Url supplied.", 400)
		return
	}

	id := storage.NextId();
	log.Println("New id for URL: ", id)

	go storage.StoreUrl(id, url.String())
	slug, err := storage.IdToSlug(id)

	if err != nil {
		http.Error(w, "Failed to create slug for URL.", 503)
		return
	}

	fmt.Fprintf(w, "http://%s:%s/%s", ServerHost, ServerPort, slug)
}

func ExpandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method - only POST allowed.", 405)
		return
	}

	slug := r.URL.Path[1:]
	log.Println("Expand Handler", slug)

	id, err := storage.SlugToId(slug)
	log.Println("Got id", id)

	if err != nil {
		http.Error(w, "Invalid slug supplied.", 400)
		return
	}

	ch := make(chan string)
	go storage.LoadUrl(id, ch)
	url := <-ch

	if len(url) == 0 {
		http.Error(w, "Failed to load url after decoding slug.", 503)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
