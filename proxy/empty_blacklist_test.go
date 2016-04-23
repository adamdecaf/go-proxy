package proxy

import (
	"net/url"
	"testing"
)

func TestEmptyBlacklist(t *testing.T) {
	bl := NewEmptyBlacklist()

	hostnames := []string{"http://localhost", "http://google.com"}

	for i := range hostnames {
		url, err := url.Parse(hostnames[i])
		if err != nil {
			t.Fatalf("error with parsing url '%s'", err)
		}

		req := Request{
			URL: *url,
			Method: GET,
		}

		e := bl.IsBlacklisted(req)
		if e != nil {
			t.Fatalf("EmptyBlacklist should never reject anything hostname - %s", req)
		}
	}
}
