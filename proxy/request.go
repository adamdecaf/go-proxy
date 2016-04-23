package proxy

import (
	"net"
	"net/url"
	"time"
)

const (
	// HTTP Methods
	GET = iota
	POST = iota
)

// Request holds onto the information from the requestor to make the
// request on their behalf. Often we want to bring in as little
// information as possible from them.
//
// There are a few options along the request to be included
// for metrics and performance monitoring.
type Request struct {
	URL url.URL
	Method int
	SourceAddress net.IP

	// private
	requestStart time.Time
}
