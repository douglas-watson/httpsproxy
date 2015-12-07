package httpsproxy

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"appengine"
	"appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", handler)
}

// Rebuild returns the full secondary HTTPS url passed as a path to the apps
// URL: for http://me.apps.com/https://ex.com/path?query=this#that, return
// https://ex.com/path?query=this#that
func Rebuild(url *url.URL) string {
	rec := strings.TrimLeft(url.Path, "/")

	if url.RawQuery != "" {
		rec += "?" + url.RawQuery
	}

	if url.Fragment != "" {
		rec += "#" + url.Fragment
	}

	return rec
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get sheet (as the unique element in URL path)
	fullURL := `https://` + Rebuild(r.URL)

	// Download the original
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(fullURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.Copy(w, resp.Body)
}
