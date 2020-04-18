package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var config = Config{
	Port:           8080,
	StubsDirectory: "./test-stubs",
}

func main() {
	err := config.LoadStubs()
	if err != nil {
		exitWithError(fmt.Sprintf("error loading test-stubs: %s", err))
	}

	serveStubs()
}

func serveStubs() {
	router := mux.NewRouter()

	for _, stub := range config.Stubs {
		fmt.Printf("Registering %s %s\n", stub.Request.Method, stub.Request.Url)
		router.HandleFunc(stub.Request.Url, func(w http.ResponseWriter, r *http.Request) {
			for key, value := range stub.Response.Headers {
				w.Header().Set(key, value)
			}
			w.WriteHeader(stub.Response.Status)

			_, err := w.Write([]byte(stub.Response.Body))
			if err != nil {
				exitWithError(fmt.Sprintf("error writing response: %s", err))
			}
		}).Methods(stub.Request.Method...)
	}

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router)
	if err != nil {
		exitWithError(fmt.Sprintf("error starting server: %s", err))
	}
}

func exitWithError(message string) {
	_, _ = fmt.Fprint(os.Stderr, message)
	os.Exit(1)
}