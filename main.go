package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
)

var config = Config{
	WriteTimeout: 15000 * time.Millisecond,
	ReadTimeout:  15000 * time.Millisecond,
	IdleTimeout:  60000 * time.Millisecond,
}

func main() {
	_, err := flags.Parse(&config)
	if err != nil {
		exitWithError(fmt.Sprintf("unable to parse command line flags: %s", err))
	}

	err = config.LoadStubs()
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
		fmt.Printf("Allowing CORS from origin '%s'\n", config.CorsAllowOrigin)
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

	fmt.Printf("Serving stubs on port %d\n", config.Port)
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
	printError(message)
	os.Exit(1)
}

func printError(message string) {
	_, _ = fmt.Fprint(os.Stderr, message)
}
