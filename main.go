package main

import "github.com/myugen/tdd-pairing-go/pkg/httpserver"

func main() {
	server := httpserver.Setup()
	server.Logger.Fatal(server.Start(":8080"))
}
