package proxyondemand_test

import (
	// "fmt"
	"github.com/papulovskiy/goproxyondemand/proxyondemand"
	"net/http"
	"testing"
)

func TestSimpleStart(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/")
	if resp != nil || err == nil {
		t.Fatal("Server is already running")
	}

	go proxyondemand.Start("localhost:8080")

	resp, err = http.Get("http://localhost:8080/")
	if resp == nil || err != nil {
		t.Fatal("Server is not running")
	}

	proxyondemand.Stop()
	resp, err = http.Get("http://localhost:8080/")
	if resp != nil || err == nil {
		t.Fatal("Server is still running")
	}

}
