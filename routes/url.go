package routes

import (
	"fmt"
	"github.com/adamdecaf/go-proxy/codec"
	"github.com/adamdecaf/go-proxy/proxy"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// shared proxy instance
var DefaultProxy = proxy.NewProxy()

func ProxyUrl(w http.ResponseWriter, r *http.Request) {
	decode_slug := readEncodedUrl(r)

	url, err := url.Parse(strings.TrimSpace(decode_slug))
	if err != nil {
		log.Printf("slug '%s' doesn't seem to be a valid url", decode_slug)
		respondWithError(w)
		return
	}

	// make the http call
	res, err := DefaultProxy.Get(*url, *r)
	if err != nil {
		log.Printf("error making request to url '%s' (err=%s)\n", url.String(), err)
		respondWithError(w)
		return
	}

	resp, err := ioutil.ReadAll(res.Reader)
	if err != nil {
		log.Printf("error reading response body from '%s' (err=%s)\n", url.String(), err)
		respondWithError(w)
		return
	}

	fmt.Fprintf(w, string(resp))
}

func readEncodedUrl(r *http.Request) string {
	// strip off /url/ from the path beginning
	raw_slug := strings.Replace(r.URL.Path, "/url/", "", 1)
	return codec.FromBase64(raw_slug)
}

func respondWithError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "We're sorry, but this request has failed.")
}
