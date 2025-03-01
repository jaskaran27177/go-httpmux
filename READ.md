# HTTPMux - Custom HTTP Multiplexer

## Overview
HTTPMux is a custom HTTP multiplexer written in Go, allowing you to map URL paths to handler functions. It works similarly to Go's `http.ServeMux`, but is implemented from scratch for educational purposes.

## Usage
- **Register Handlers**: Use `HandlePathFunc` to map a URL path to a handler function.
- **Serve Requests**: Pass the `Mux` instance to `http.ListenAndServe()` to handle incoming requests.

## Example
```go
package main

import (
	"fmt"
	"net/http"
	httpsrv "myproject/srv/http"
)

func main() {
	httpsrv.Mux.HandlePathFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from HTTPMux!")
	}, "/hello")

	http.ListenAndServe(":8080", &httpsrv.Mux)
}
