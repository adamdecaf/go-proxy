package proxy

import (
	"errors"
	"github.com/hashicorp/go-retryablehttp"
	"log"
	"net"
	"net/http"
	"net/url"
)

var (
	// defaults
	DefaultHttpClient = retryablehttp.NewClient()

	// errors
	UnknownMethodError = errors.New("Unknown method for proxy request")
)

// A Proxy is a shared interface for making requests it is supposed
// to be shared across goroutines and also will perform changes to
// the request and response in accordance to this proxy's requirements
// (headers, response parsing / transforms, etc).
//
// Usage
//  var DefaultProxy := NewProxy()
//
//  res, err := DefaultProxy.Get("http://example.com")
//
type Proxy struct {
	client *retryablehttp.Client
}

// NewProxy() creates a new instance of the proxy to be
// shared across goroutines
func NewProxy() Proxy {
	proxy := Proxy{
		client: DefaultHttpClient,
	}
	return proxy
}

func (p Proxy) Get(url url.URL, r http.Request) (*Response, error) {
	req := Request{
		URL: url,
		Method: GET,
		SourceAddress: parseSourceAddress(r),
	}
	return request(req)
}

func get(req Request) (*Response, error) {
	r, err := retryablehttp.Get(req.URL.String())
	if err != nil {
		return nil, err
	}

	resp := Response{
		Reader: r.Body,
	}

	return &resp, nil
}

// `request` checks the blacklist and rejects requests without performing them
// if the blacklist is triggered.
func request(req Request) (*Response, error) {
	if req.SourceAddress == nil {
		log.Printf("no source address on request for '%s'", req.URL.String())
		return nil, SourceAddressBlacklisted
	}

	err := DefaultBlacklist.IsBlacklisted(req)
	if err != nil  {
		return nil, *err
	}

	r, e := makeRequest(req)
	if e != nil {
		return nil, e
	}

	// fold over transformers off the original response
	if r != nil {
		for i := range DefaultTransformers {
			morphed := DefaultTransformers[i].Transform(*r)
			r = &morphed
		}
	}

	return r, nil
}

// `makeRequest` just gives us a wrapper around returning tuples of
// request execution.
func makeRequest(req Request) (*Response, error) {
	switch req.Method {
	default:
		return nil, UnknownMethodError
	case GET:
		return get(req)
	}
}

func parseSourceAddress(r http.Request) net.IP {
	if h := r.Header.Get("X-Real-Ip"); h != "" {
                return parseFoundIP(h)
        }

	if h := r.Header.Get("X-Forwarded-For"); h != "" {
		return parseFoundIP(h)
	}

	if h := r.Header.Get("Remote-Address"); h != "" {
		return parseFoundIP(h)
	}

	return nil
}

func parseFoundIP(ip string) net.IP {
	return net.ParseIP(ip)
}
