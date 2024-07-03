package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	println("Content-Encoding: ", r.Header.Get("Content-Encoding"))
	println("Content-Type: ", r.Header.Get("Content-Type"))
	println("Content-Disposition: ", r.Header.Get("Content-Disposition"))

	buf, err := gzip.NewReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	all, err := io.ReadAll(buf)
	if err != nil {
		log.Fatal(err)
	}

	println(string(all))
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
