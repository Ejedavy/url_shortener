package urlshortner

import "net/http"

func MapHandler(pathToURL map[string]string, fallbackHandler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if actual, found := pathToURL[path]; found {
			http.Redirect(w, r, actual, http.StatusFound)
			return
		} else {
			fallbackHandler.ServeHTTP(w, r)
		}
	}
}

func YAMLHandler() {

}
