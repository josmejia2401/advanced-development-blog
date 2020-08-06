package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const APP_NAME string = "MS_REVERSE_PROXY"

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Log.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(HandleRequestAndRedirect)
	mux.Handle("/", middlewareOne(finalHandler))

	srvHTTPs := &http.Server{
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":4443",
		Handler:      mux,
	}
	Log.Fatal(srvHTTPs.ListenAndServe())
}
