package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var config = Config{
	Port:           8080,
	StubsDirectory: "./test-stubs",
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
		router.HandleFunc(stub.Request.Url, StubHandler(stub)).MatcherFunc(StubMatcher(stub))
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("0.0.0.0:%d", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}
	err := srv.ListenAndServe()
	if err != nil {
		exitWithError(fmt.Sprintf("error starting server: %s", err))
	}
}

func exitWithError(message string) {
	_, _ = fmt.Fprint(os.Stderr, message)
	os.Exit(1)
}
