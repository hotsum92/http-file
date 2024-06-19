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

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("content-type: ", r.Header.Get("Content-Type"))

	unzip(buf)
}

func unzip(file []byte) {
	// zipファイルを読み込む
	r, err := zip.NewReader(bytes.NewReader(file), int64(len(file)))
	if err != nil {
		log.Fatal(err)
	}

	// 各ファイルを展開する
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()

		// ファイル名を出力する
		fmt.Println(f.Name)

		// ファイルを展開する
		buf, err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}

		// ファイルの内容を出力する
		fmt.Println(string(buf))
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
