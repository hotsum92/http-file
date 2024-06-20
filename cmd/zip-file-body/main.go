package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	println("Content-Encoding: ", r.Header.Get("Content-Encoding"))
	println("Content-Type: ", r.Header.Get("Content-Type"))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	reader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()

		fmt.Println("fileName: ", f.Name)

		buf, err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("fileContent: ", string(buf))
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
