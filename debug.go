// +build debug

package sg

import (
	"net/http"
	"net/http/httputil"
)

func init() {
	dumpRequest = func(t Tracer, request *http.Request) {
		if t == nil {
			return
		}

		if dump, err := httputil.DumpRequest(request, true); err == nil {
			t.Printf("\n%s\n", dump)
		}
	}

	dumpResponse = func(t Tracer, response *http.Response) {
		if t == nil {
			return
		}

		if dump, err := httputil.DumpResponse(response, true); err == nil {
			t.Printf("\n%s\n", dump)
		}
	}
}
