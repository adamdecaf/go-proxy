package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	proxy "github.com/adamdecaf/go-proxy/proxy"
)

// GET /url/<base64-url>
//  - extract and parse url
//  - filter from blacklist
//  - load from remote via http client
//    - need to replace headers
//    - and do all sorts of other proxy related things
//  - ship response back
//  - detect html & replace nested urls

func ProxyUrl(w http.ResponseWriter, r *http.Request) {
	proxy := proxy.NewHttpProxy()

	raw_url := "http://google.com"
	url, err := url.Parse(raw_url)
	if err != nil {
		log.Printf("error parsing url '%s'", raw_url)
		// todo: BackRequest
	}

	resp, err := proxy.Request(*url)

	if err != nil {
		log.Printf("error making request to %s", url)
		// todo: BackRequest
	}

	fmt.Println(resp)

	fmt.Fprintf(w, "ProxyUrl")
}
