package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"facts_and_questions_api/handler"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	hs := newHTTPServer(getAddr(), handler.NewHandler(handler.LogWith(logger)))

	logger.Printf("Listening on http://127.0.0.1%s\n", hs.Addr)
	hs.ListenAndServe()
}

func newHTTPServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func getAddr() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}

	return ":8383"
}
