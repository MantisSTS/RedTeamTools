package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	socks5 "github.com/things-go/go-socks5"
)

func main() {

	// Get flags for IP and port
	ip := flag.String("ip", "0.0.0.0", "IP address to listen on")
	port := flag.Int("port", 8888, "Port to listen on")
	flag.Parse()

	log.Printf("Starting server on: %s:%d", *ip, *port)

	// Create a SOCKS5 server
	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
	)

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", fmt.Sprintf("%s:%d", *ip, *port)); err != nil {
		panic(err)
	}
}
