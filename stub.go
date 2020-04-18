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

func QueryMatcher(query KeyValuePairs) func(*http.Request, *mux.RouteMatch) bool {
	return func(r *http.Request, rm *mux.RouteMatch) bool {
		urlQuery := r.URL.Query()
		for key, value := range query {
			if matched, _ := regexp.MatchString(value, urlQuery.Get(key)); !matched {
				return false
			}
		}

		return true
	}
}

func HeadersMatcher(headers KeyValuePairs) func(*http.Request, *mux.RouteMatch) bool {
	return func(r *http.Request, rm *mux.RouteMatch) bool {
		for key, value := range headers {
			if matched, _ := regexp.MatchString(value, r.Header.Get(key)); !matched {
				return false
			}
		}

		return true
	}
}