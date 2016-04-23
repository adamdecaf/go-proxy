package proxy

import (
	"net"
	"net/url"
	"testing"
)

func TestJSONBlacklistHostnames(t *testing.T) {
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

func TestJSONBlacklistIPsAndSubnets(t *testing.T) {
	bl, err := NewJSONBlacklistFromFile("./blacklist.json") // paths are funky
	if err != nil {
		t.Fatalf("error creating json blacklist err=%s", err)
	}

	// check source ip
	ip1 := net.ParseIP("127.0.0.1")
	req1 := Request{
		SourceAddress: ip1,
	}
	e1 := bl.IsBlacklisted(req1)
	if e1 == nil {
		t.Fatalf("request from '127.0.0.1' should have been blocked")
	}

	// check subnet
	ip2 := net.ParseIP("10.2.3.5")
	req2 := Request{
		SourceAddress: ip2,
	}
	e2 := bl.IsBlacklisted(req2)
	if e2 == nil {
		t.Fatalf("request from '10.2.3.5' should have been blocked")
	}
}
