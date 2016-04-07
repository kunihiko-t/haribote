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
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandler(t *testing.T) {

	config, err := ioutil.ReadFile("config_sample.json")
	if err != nil {
		t.Error("Cannot Read file", err)
		return
	}

	var c JsonConfig
	err = json.Unmarshal(config, &c)

	if err != nil {
		t.Error(err)
		return
	}

	for _, s := range c.Server {
		h := &MockHandler{Method: s.Method, ContentType: s.ContentType, StatusCode: s.StatusCode, Text: s.Text, File: s.File}
		ts := httptest.NewServer(h)
		defer ts.Close()
		var err []error
		var res *http.Response
		request := gorequest.New()
		switch h.Method {
		case "GET":
			res, _, err = request.Get(ts.URL).End()
		case "POST":
			res, _, err = request.Post(ts.URL).End()
		case "PUT":
			res, _, err = request.Put(ts.URL).End()
		case "DELETE":
			res, _, err = request.Delete(ts.URL).End()
		default:
			t.Error("unexpected method")
			return
		}

		if err != nil {
			t.Error("unexpected error:", err)
		}

		if h.StatusCode != 0 && res.StatusCode != h.StatusCode {
			t.Error("unexpected status code")
		}

		if h.ContentType != res.Header.Get("Content-Type") {
			t.Error("unexpected content type")
		}

		if h.File != "" {

			expectedFile, err := ioutil.ReadFile(h.File)
			if err != nil {
				t.Error("Cannot Read file", err)
			}

			respBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Error("Cannot Read response", err)
			}

			if reflect.DeepEqual(respBody, expectedFile) != true {
				t.Error("Unexpected Response", err)
			}

		} else if h.Text != "" {

			respBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Error("Cannot Read response", err)
			}

			if string(respBody) != h.Text {
				t.Error("Unexpected Response", err)
			}

		}

	}

}
