package proxy

import (
	"net/url"
	"testing"
)

func TestJSONBlacklist(t *testing.T) {
	bl, err := NewJSONBlacklistFromFile("./blacklist.json") // paths are funky
	if err != nil {
		t.Fatalf("error creating json blacklist err=%s", err)
	}

	blocked := []string{"http://localhost"}
	unblocked := []string{"http://google.com"}

	// blocked

	for i := range blocked {
		url, err := url.Parse(blocked[i])
		if err != nil {
			t.Fatalf("error with parsing url '%s'", err)
		}

		req := Request{
			URL: *url,
			Method: GET,
		}

		e := bl.IsBlacklisted(req)
		if e == nil {
			t.Fatalf("JSONBlacklist should reject this hostname %s - %s", blocked[i], req.URL.Host)
		}
	}

	// unblocked

	for i := range unblocked {
		url, err := url.Parse(unblocked[i])
		if err != nil {
			t.Fatalf("error with parsing url '%s'", err)
		}

		req := Request{
			URL: *url,
			Method: GET,
		}

		e := bl.IsBlacklisted(req)
		if e != nil {
			t.Fatalf("JSONBlacklist should not reject this hostname %s - %s", unblocked[i], req.URL.Host)
		}
	}
}
