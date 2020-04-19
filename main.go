package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var config = Config{
	Port:            8080,
	WriteTimeout:    15 * time.Second,
	ReadTimeout:     15 * time.Second,
	IdleTimeout:     60 * time.Second,
	StubsDirectory:  "./test-stubs",
	CorsAllowOrigin: "*",
}

func main() {
	err := config.LoadStubs()
	if err != nil {
		exitWithError(fmt.Sprintf("error loading stubs: %s", err))
	}

	serveStubs()
}

func serveStubs() {
	router := mux.NewRouter()

	for _, stub := range config.Stubs {
		fmt.Printf("Registering %s\n", stub.String())
		if config.IsCorsEnabled() && isMissingOptionsMethod(stub.Request.Method) {
			stub.Request.Method = append(stub.Request.Method, http.MethodOptions)
		}
		router.HandleFunc(stub.Request.Url, StubHandler(stub)).
			Methods(stub.Request.Method...).
			MatcherFunc(QueryMatcher(stub.Request.Query)).
			MatcherFunc(HeadersMatcher(stub.Request.Headers))
	}

	if config.IsCorsEnabled() {
		router.Use(mux.CORSMethodMiddleware(router))
		router.Use(CORSBlanketMiddleware)
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("0.0.0.0:%d", config.Port),
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		IdleTimeout:  config.IdleTimeout,
	}
	err := srv.ListenAndServe()
	if err != nil {
		exitWithError(fmt.Sprintf("error starting server: %s", err))
	}
}

func isMissingOptionsMethod(methods []string) bool {
	for _, method := range methods {
		if method == http.MethodOptions {
			return false
		}
	}
	return true
}

func exitWithError(message string) {
	_, _ = fmt.Fprint(os.Stderr, message)
	os.Exit(1)
}
