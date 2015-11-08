package proxyondemand

import (
	"github.com/elazarl/goproxy"
	"net/http"
	// "net/url"
	"os"
	// "os/exec"
	"fmt"
	"html"
	"log"
	"sync"
)

var MinPort, MaxPort uint
var Ports = struct {
	sync.RWMutex
	m map[uint]bool
}{m: make(map[uint]bool)}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Stop() {
	log.Println("Shutdown requested")
	os.Exit(0)
}

func GetNextAvailablePort() uint {
	var p uint
	for p = MinPort; p <= MaxPort; p++ {
		// fmt.Printf("%+v %+v\n", p, Ports[p])
		if !Ports.m[p] {
			return p
		}
	}
	return 0
}

func CreateProxy() (uint, error) {
	Ports.Lock()
	Port := GetNextAvailablePort()
	Ports.m[Port] = true
	Ports.Unlock()
	log.Printf("Proxy create request, port: %d", Port)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	go http.ListenAndServe(fmt.Sprintf(":%d", Port), proxy)
	return Port, nil
}

func Start(bind string) {
	// TODO: implement command line arguments via flag
	MinPort = 40000
	MaxPort = 45000
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
