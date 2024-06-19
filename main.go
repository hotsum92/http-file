package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//dump, err := httputil.DumpRequest(r, true)

	//if err != nil {
	//	http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	//	return
	//}

	mr, _ := r.MultipartReader()
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}

		fmt.Println(p.FormName(), ":", p.FileName(), p.Header.Get("Content-Type"))
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
