/*
The MIT License (MIT)

Copyright (c) 2016 kunihiko-t.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type ServerParam struct {
	Path        string
	Method      string
	ContentType string
	StatusCode  int
	Text        string
	File        string
}

type JsonConfig struct {
	Server []ServerParam
}

type MockHandler struct {
	Method      string
	ContentType string
	StatusCode  int
	Text        string
	File        string
}

func (h *MockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != strings.ToUpper(h.Method) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", h.ContentType)
	if h.StatusCode != 0 {
		w.WriteHeader(h.StatusCode)
	}

	if h.File != "" {
		body, err := ioutil.ReadFile(h.File)
		checkError(err)
		w.Write(body)
	} else {
		fmt.Fprintf(w, h.Text)
	}

}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, " Usage of Haribote: \nharibote [OPTIONS] ARGS...\nOptions\n")
		flag.PrintDefaults()
	}

	port := flag.Int("p", 9090, "Port number")
	fileName := flag.String("f", "", "Config file name")
	flag.Parse()

	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	config, err := ioutil.ReadFile(*fileName)
	checkError(err)
	var c JsonConfig
	err = json.Unmarshal(config, &c)
	checkError(err)

	for _, s := range c.Server {
		h := &MockHandler{Method: s.Method, ContentType: s.ContentType, StatusCode: s.StatusCode, Text: s.Text, File: s.File}
		http.Handle(s.Path, h)
		log.Printf("Registered Handler %v", s)
	}

	addr := fmt.Sprintf(":%d", *port)
	err = http.ListenAndServe(addr, nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
