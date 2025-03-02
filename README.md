# HTTPMux - Custom HTTP Multiplexer

## Overview

HTTPMux is a custom HTTP multiplexer written in Go, allowing you to map URL paths to handler functions. It works similarly to Go's `http.ServeMux`, but is implemented from scratch for educational purposes.

## Documentation

https://pkg.go.dev/github.com/jaskaran27177/go-httpmux

## Usage

- **Register Handlers**: Use `HandlePathFunc` to map a URL path to a handler function.
- **Serve Requests**: Pass the `Mux` instance to `http.ListenAndServe()` to handle incoming requests.

## Example

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/jaskaran27177/go-httpmux"
)

// Declare the global Mux variable
var Mux httpmux.HTTPMux

func main() {
	// Register a handler using HandlePathFunc
	Mux.HandlePathFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from HTTPMux!")
	}, "/hello")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", &Mux)
}

```
