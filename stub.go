package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func StubHandler(stub Stub) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for key, value := range stub.Response.Headers {
			w.Header().Set(key, value)
		}
		w.WriteHeader(stub.Response.Status)

		_, err := w.Write([]byte(stub.Response.Body))
		if err != nil {
			exitWithError(fmt.Sprintf("error writing response: %s", err))
		}
	}
}

func StubMatcher(stub Stub) func(*http.Request, *mux.RouteMatch) bool {
	return func(r *http.Request, rm *mux.RouteMatch) bool {
		if matched := matchMethod(stub.Request.Method, r.Method); !matched {
			return false
		}

		for key, value := range stub.Request.Headers {
			if matched, _ := regexp.MatchString(value, r.Header.Get(key)); !matched {
				return false
			}
		}

		query := r.URL.Query()
		for key, value := range stub.Request.Query {
			if matched, _ := regexp.MatchString(value, query.Get(key)); !matched {
				return false
			}
		}

		return true
	}
}

func matchMethod(methods []string, method string) bool {
	for _, allowedMethod := range methods {
		if method == allowedMethod {
			return true
		}
	}
	return false
}
