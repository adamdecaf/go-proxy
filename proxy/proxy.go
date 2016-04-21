// package proxy

// import (
// 	"fmt"
// 	"io"
// 	"net/url"
// 	"os"
// )

// // reading
// // golang.org/x/net/proxy
// // github.com/elazarl/goproxy
// // PHProxy script (in Dropbox)

// type Response interface {
// 	Reader() io.Reader
// }

// // todo: EmptyResponse (for errors)

// type FullResponse struct {
// 	Response
// }

// func (r FullResponse) Reader() io.Reader {
// 	reader, _ := os.Open("/dev/null")

// 	// if err != nil {
// 	// 	return "", err
// 	// }

// 	// r := io.NewReader(reader)
// 	return reader
// }

// type Proxy interface {
// 	Request(url url.URL) (Response, error)
// }

// type HttpProxy struct {
// 	Proxy
// }

// func (p HttpProxy) Request(url url.URL) (Response, error) {
// 	blacklist, err := BlacklistFromFile()
// 	if err != nil {
// 		// log.Printf("")
// 		return FullResponse{}, err // todo: EmptyResponse
// 	}

// 	if blacklist.Contains(url.Host) {
// 		err := fmt.Errorf("url '%s' is in blacklist, ignoring", url)
// 		return FullResponse{}, err // todo: EmptyResponse
// 	}

// 	return FullResponse{}, nil
// }

// func NewHttpProxy() Proxy {
// 	return HttpProxy{}
// }
