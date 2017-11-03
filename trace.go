package sg

import "net/http"

// Tracer is the implementation of a tracer that can print debug information of
// a Client SendGrid API requests and responses. Setting a tracer on the client
// can help us debug errors.
//
// The standard Logger implements this interface.
type Tracer interface {
	Printf(format string, v ...interface{})
}

var dumpRequest func(Tracer, *http.Request)
var dumpResponse func(Tracer, *http.Response)
