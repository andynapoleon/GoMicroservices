package main

import (
	"log"

	"github.com/andynapoleon/GoDistributedService/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080") // "server" here is the go file name in the package imported (the last one)
	log.Fatal(srv.ListenAndServe())
}
