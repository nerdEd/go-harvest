package harvest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func mockHTTPClient(code int, body string) (s *httptest.Server, h *http.Client) {

	// Test server that always responds with 200 code, and specific payload
	s = httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, body)
		},
	))

	// Make a transport that reroutes all traffic to the example server
	transport := &http.Transport{
		DisableCompression: true,
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(s.URL)
		},
	}

	// Make a http.Client with the transport
	h = &http.Client{Transport: transport}
	return
}
