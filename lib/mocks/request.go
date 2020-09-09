package mocks

import (
	"fmt"
	"net/http"
)

// Responders are callbacks that receive and http request and return a mocked response.
type Responder func(*http.Request) (*http.Response, error)

// MockTransport implements http.RoundTripper, which fulfills single http requests issued by
// an http.Client.  This implementation doesn't actually make the call, instead defering to
// the registered list of responders.
type MockTransport struct {
	responders      map[string]Responder
}

// RoundTrip is required to implement http.MockTransport.  Instead of fulfilling the given request,
// the internal list of responders is consulted to handle the request.  If no responder is found
// an error is returned, which is the equivalent of a network error.
func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	key := req.Method + " " + req.URL.String()

	// scan through the responders and find one that matches our key
	for k, r := range m.responders {
		if k != key {
			continue
		}
		return r(req)
	}

	return nil, fmt.Errorf("no responder speficied")
}

// RegisterResponder adds a new responder, associated with a given HTTP method and URL.  When a
// request comes in that matches, the responder will be called and the response returned to the client.
func (m *MockTransport) RegisterResponder(method, url string, responder Responder) {
	if m.responders == nil {
		m.responders = make(map[string]Responder)
	}

	m.responders[method+" "+url] = responder
}
