package main

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	println("Content-Encoding: ", r.Header.Get("Content-Encoding"))
	println("Content-Type: ", r.Header.Get("Content-Type"))

	reader, err := gzip.NewReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	println(string(buf))
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
