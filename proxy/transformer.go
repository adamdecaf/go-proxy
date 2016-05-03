package proxy

import (
	"net/url"
)

// Transformer is a type which performs some action on a `Response`
// and returns another `Response`.
//
// Typical use cases are for doing things with html or metrics.
type Transformer interface {
	Transform(url.URL, Response) Response
}
