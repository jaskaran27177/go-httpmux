package httpmux

import (
	"errors"
	"net/http"
	"sync"
)

// HTTPMux is a custom HTTP multiplexer that maps URL paths to handlers.
// It includes a map to store handlers and a mutex to ensure thread safety.
type HTTPMux struct {
	pathHandlers map[string]http.Handler
	mu           sync.Mutex
}

// ServeHTTP implements the http.Handler interface for HTTPMux.
// It looks up the request path in the pathHandlers map and forwards the request to the corresponding handler.
// If no handler is found, it returns a 404 Not Found response.
// This method is safe for concurrent use.
func (receiver *HTTPMux) ServeHTTP(respwriter http.ResponseWriter, req *http.Request) {
	receiver.mu.Lock()
	defer receiver.mu.Unlock()

	// Lazily initialize the map if needed
	if receiver.pathHandlers == nil {
		receiver.pathHandlers = make(map[string]http.Handler)
	}

	if handler, ok := receiver.pathHandlers[req.URL.Path]; ok {
		handler.ServeHTTP(respwriter, req)
		return
	}
	http.NotFound(respwriter, req)
}

// HandlePath registers an http.Handler for the specified path.
// It returns an error if a handler is already registered for that path.
// This method is safe for concurrent use.
func (receiver *HTTPMux) HandlePath(handler http.Handler, path string) error {
	receiver.mu.Lock()
	defer receiver.mu.Unlock()

	// Lazily initialize the map if needed
	if receiver.pathHandlers == nil {
		receiver.pathHandlers = make(map[string]http.Handler)
	}

	if receiver.pathHandlers[path] != nil {
		return errors.New("path already exists")
	}
	receiver.pathHandlers[path] = handler
	return nil
}

// HandlePathFunc registers a handler function for the specified path.
// It converts the handler function into an http.HandlerFunc and registers it using HandlePath.
// This method is safe for concurrent use.
func (receiver *HTTPMux) HandlePathFunc(fn func(http.ResponseWriter, *http.Request), path string) {
	receiver.HandlePath(http.HandlerFunc(fn), path)
}

