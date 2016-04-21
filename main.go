package main

import (
	"fmt"
	"flag"
	"net/http"
	"log"
	"github.com/adamdecaf/go-proxy/routes"
)

const DefaultHttpPort = 8888

func main() {
	// port overrride
	var portFromFlag int
	help := fmt.Sprintf("pick a port other than the default (%d)", DefaultHttpPort)
	flag.IntVar(&portFromFlag, "p", DefaultHttpPort, help)

	if portFromFlag <= 0 || portFromFlag > 65535 {
		log.Fatalf(fmt.Sprintf("invalid port given %d", portFromFlag))
	}

	flag.Parse()

	// Setup and start the http server
	log.Printf("Starting http server on :%d\n", portFromFlag)

	err := startHttpServer(portFromFlag)
	if err != nil {
		header := fmt.Sprintf("error binding to port %d\n", portFromFlag)
		log.Fatalf(header, err)
	}
}

func startHttpServer(port int) *error {
	http.Handle("/", http.FileServer(http.Dir("./html/")))
	http.HandleFunc("/ping", routes.Ping)
	http.HandleFunc("/url/", routes.ProxyUrl)

	listen := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(listen, nil)

	return &err
}
