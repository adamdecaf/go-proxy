package proxy

import (
	"time"
)

// Request holds onto the information from the requestor to
// make the request on their behalf. Often we want to bring
// in as little information as possible from them.
//
// There are a few options along the request to be included
// for metrics and performance monitoring.
type Request struct {
	URL string
	Method int

	// private
	requestStart time.Time
}

const (
	GET = iota
)

// todo
// func IsBlacklisted() boolean

// could be from:
// - source ip (or range of ips)
// - url, host, query params, etc

// blacklist, err := BlacklistFromFile()
// if err != nil {
// 	// log.Printf("")
// 	return FullResponse{}, err // todo: EmptyResponse
// }

// if blacklist.Contains(url.Host) {
// 	err := fmt.Errorf("url '%s' is in blacklist, ignoring", url)
// 	return FullResponse{}, err // todo: EmptyResponse
// }
