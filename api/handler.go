// This package is based on this great article about error handling with golang headers
//   https://elithrar.github.io/article/http-handler-error-handling-revisited/

package api

import (
	"log"
	"net/http"
)

// Env holds application-wide configuration.
type Env struct {
	Secret string
}

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	H func(w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
// See:
//    https://annevankesteren.nl/2015/02/same-origin-policy
//    http://stackoverflow.com/questions/22972066/how-to-handle-preflight-cors-requests-on-a-go-server
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "x-authentication")

	err := h.H(w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			//http.Error(w, http.StatusText(http.StatusInternalServerError),
			//	http.StatusInternalServerError)
		}
	}
}
