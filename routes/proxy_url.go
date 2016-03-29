package routes

import (
	"fmt"
	"net/http"
)

// GET /url/<base64-url>
//  - parse url
//  - filter from blacklist
//  - load
//  - detect html & replace nested urls

func ProxyUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ProxyUrl")
}
