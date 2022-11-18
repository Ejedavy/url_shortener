package main

import (
	"fmt"
	"net/http"

	urlshortner "github.com/ejedavy/url_shortner"
)

func main() {
	mux := createDefaultMultiplexer()
	var mapOfURLS map[string]string
	mapOfURLS = generateTestMap()
	handler := urlshortner.MapHandler(mapOfURLS, mux)
	fmt.Println("Starting the server at :8000")
	http.ListenAndServe(":8000", handler)
}

func createDefaultMultiplexer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("Hello!\nYou are seeing this page because the link you requested is not available on the server"))
	})
	return mux
}

func generateTestMap() map[string]string {
	mapOfURLS := map[string]string{
		"/test":  "http://www.goal.com",
		"/test2": "http://www.google.com",
	}
	return mapOfURLS
}
