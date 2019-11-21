package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var CommitSHA string

var (
	wwwRoot = flag.String("root", "/usr/share/booktaxi", "path to serve files from")
	port    = flag.Int("port", 8080, "port to listen on")
)

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(*wwwRoot))
	http.Handle("/", fs)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("version %s listening on %s\n", CommitSHA, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
