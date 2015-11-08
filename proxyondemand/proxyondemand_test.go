package proxyondemand_test

import (
	"encoding/json"
	// "fmt"
	"github.com/papulovskiy/goproxyondemand/proxyondemand"
	"io/ioutil"
	"net/http"
	"testing"
)

type ProxyDescription struct {
	Port uint `json:"port"`
}

type ProxyRequest struct {
	Url          string
	Status       uint
	ResponseTime float32
}

type ProxyLog struct {
	Requests []ProxyRequest
}

// TODO: fix the server start and shutdown
func TestSimpleStartStop(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/")
	if resp != nil || err == nil {
		t.Fatal("Server is already running")
	}

	go proxyondemand.Start("localhost:8080")

	resp, err = http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal("Server is not running")
	}

	// proxyondemand.Stop()
	// resp, err = http.Get("http://localhost:8080/")
	// if resp != nil || err == nil {
	// 	t.Fatal("Server is still running")
	// }
}

func TestSimpleProxyCreation(t *testing.T) {
	resp, err := http.PostForm("http://localhost:8080/create", nil)
	if resp == nil || err != nil {
		t.Fatal("Cannot create a proxy")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Cannot read response from server")
	}

	var port ProxyDescription
	err = json.Unmarshal(body, &port)
	if err != nil {
		t.Fatal("Cannot parse response from server")
	}

	if port.Port == 0 {
		t.Fatal("Port cannot be equal zero")
	}

}
