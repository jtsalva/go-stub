package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/gorilla/mux"
)

func StubHandler(stub Stub) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		for key, value := range stub.Response.Headers {
			w.Header().Set(key, value)
		}
		w.WriteHeader(stub.Response.Status)

		response := []byte(stub.Response.Body)
		if stub.Response.File != "" {
			response, err = ioutil.ReadFile(stub.Response.File)
			if err != nil {
				printError(fmt.Sprintf("[%s %s] error reading file '%s': %s", r.Method, stub.Request.Path, stub.Response.File, err))
			}
		}
		_, err = w.Write(response)
		if err != nil {
			printError(fmt.Sprintf("error writing response: %s", err))
		}

		if stub.Response.Latency != 0 {
			time.Sleep(time.Duration(stub.Response.Latency) * time.Millisecond)
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
