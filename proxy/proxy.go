package proxy

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

// reading
// golang.org/x/net/proxy
// github.com/elazarl/goproxy
// PHProxy script (in Dropbox)

type Response interface {
	Reader() bufio.Reader
}

// todo: EmptyResponse (for errors)

type FullResponse struct {
	Response
}

func (r FullResponse) Reader() bufio.Reader {
	reader, _ := os.Open("/dev/null")
	// if err != nil {
	// 	return "", err
	// }

	buf := bufio.NewReader(reader)
	return *buf
}

type Proxy interface {
	Request(url url.URL) (Response, error)
}

type HttpProxy struct {
	Proxy
}

func (p HttpProxy) Request(url url.URL) (Response, error) {
	blacklist, err := BlacklistFromFile()
	if err != nil {
		// log.Printf("")
		return FullResponse{}, err // todo: EmptyResponse
	}

	if blacklist.Contains(url.Host) {
		err := fmt.Errorf("url '%s' is in blacklist, ignoring", url)
		return FullResponse{}, err // todo: EmptyResponse
	}

	return FullResponse{}, nil
}

func NewHttpProxy() Proxy {
	return HttpProxy{}
}
