package routes

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"net/url"
	"io/ioutil"
	// "github.com/adamdecaf/go-proxy/proxy"
)

// GET /url/<base64-url>
//  - extract and parse url
//  - filter from blacklist
//  - load from remote via http client
//    - need to replace headers
//    - and do all sorts of other proxy related things
//  - ship response back
//  - detect html & replace nested urls

// use this
const MaxUrlLength = 1000

func ProxyUrl(w http.ResponseWriter, r *http.Request) {
	decode_slug, err := readEncodedUrl(r)
	if err != nil {
		log.Printf("error parsing path '%s' into raw slug", r.URL.Path)
		respondWithError(w)
		return
	}

	url, err := url.Parse(strings.TrimSpace(decode_slug))
	if err != nil {
		log.Printf("slug '%s' doesn't seem to be a valid url", decode_slug)
		respondWithError(w)
		return
	}

	// make the http call
	// todo: move this stuff out into the proxy folder
	//       and look at using a pooling lib / writing my own
	// todo: check blacklist and reject
	// todo: stats on hostname, response timings, etc

	client := &http.Client{}
	res, err := client.Get(url.String())
	if err != nil {
		log.Printf("error making request to url '%s' (err=%s)\n", url.String(), err)
		respondWithError(w)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading response body from '%s' (err=%s)\n", url.String(), err)
		respondWithError(w)
		return
	}

	fmt.Fprintf(w, string(body))
}

func readEncodedUrl(r *http.Request) (string, error) {
	// strip off /url/ from the path beginning
	raw_slug := strings.Replace(r.URL.Path, "/url/", "", 1)

	// attempt to decode base64
	bytes, err := base64.StdEncoding.DecodeString(raw_slug)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// todo: non-200 status code
func respondWithError(w http.ResponseWriter) {
	fmt.Fprintf(w, "We're sorry, but this request has failed.")
}
