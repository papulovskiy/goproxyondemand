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

func CreateProxy() (uint, error) {
	log.Println("Proxy create request")
	return 8081, nil
}

func Start(bind string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HandleIndex(w, r)
	})
	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		Stop()
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			port, _ := CreateProxy()
			w.Write([]byte(fmt.Sprintf("{\"port\":%d}", port)))
		default:
			log.Println("Create proxy non-POST request")
		}

	})

	log.Fatal(http.ListenAndServe(bind, nil))
}
