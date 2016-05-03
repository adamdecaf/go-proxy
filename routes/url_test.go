package routes

import (
	"net/url"
	"testing"
)

func TestAddSchemeToURL(t *testing.T) {
	u := "example.com"
	parsed, err := url.Parse(u)
	if err != nil {
		t.Fatalf("failed to parse '%s', err=%s", u, err)
	}

	after := appendMissingHTTPScheme(*parsed)

	if after.Scheme != "http" {
		t.Fatalf("didn't add scheme onto %s", after)
	}
}

func TestLeaveURLWithSchemeAlone(t *testing.T) {
	u := "http://example.com"
	parsed, err := url.Parse(u)
	if err != nil {
		t.Fatalf("failed to parse '%s', err=%s", u, err)
	}

	before := parsed
	appendMissingHTTPScheme(*parsed)

	if before != parsed {
		t.Fatalf("something changed, before=%s, parsed=%s", before, parsed)
	}
}
