package proxy

import (
	"errors"
	"github.com/hashicorp/go-retryablehttp"
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

func (p Proxy) Get(url string) (*Response, error) {
	req := Request{
		URL: url,
		Method: GET,
	}
	return request(req)
}

func get(req Request) (*Response, error) {
	r, err := retryablehttp.Get(req.URL)
	if err != nil {
		return nil, err
	}

	resp := Response{
		Reader: r.Body,
	}

	return &resp, nil
}

// todo: check blacklist w/ `if req.IsBlacklisted()`
func request(req Request) (*Response, error) {
	switch req.Method {
	default:
		return nil, UnknownMethodError
	case GET:
		return get(req)
	}
}
