package proxyondemand

import (
	// "github.com/elazarl/goproxy"
	"net/http"
	// "net/url"
	"os"
	// "os/exec"
	"fmt"
	"html"
	"log"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Stop() {
	log.Println("Shutdown requested")
	os.Exit(0)
}

func Start(bind string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HandleIndex(w, r)
	})
	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		Stop()
	})
	log.Fatal(http.ListenAndServe(bind, nil))
}
