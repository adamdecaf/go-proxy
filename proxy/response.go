package proxy

import (
	"io"
)

// Response contains a buffered output of the interaction
// performed on behalf of the requestor.e
type Response struct {
	Reader io.Reader
}
