package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"adampayne315/fetch/api"

	"github.com/gin-gonic/gin"
)

func NewGinServer(f *api.StrictFetchApi, port string) *http.Server {
	// This is how you set up a basic gin router
	r := gin.Default()
	si := api.NewStrictHandler(f, make([]api.StrictMiddlewareFunc, 0))
	api.RegisterHandlers(r, si)
	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}
	return s
}

func main() {
	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()
	// Create an instance of our handler which satisfies the generated interface
	f := api.NewStrictFetchApi()
	s := NewGinServer(f, *port)
	// run until we hit Ctrl-C
	log.Fatal(s.ListenAndServe())
}
