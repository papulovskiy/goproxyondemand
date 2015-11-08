package main

import (
	proxyondemand "github.com/papulovskiy/goproxyondemand/proxyondemand"
)

func main() {
	proxyondemand.Start("localhost:8080")
}
