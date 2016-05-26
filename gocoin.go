package main

import (
	"net/http"
	"fmt"
)
type data struct {
	to string
	from string
	amount string
}

func main() {
	http.HandleFunc("/", serveResponse)
	http.ListenAndServe(":8080", nil)
}

func serveResponse(d http.ResponseWriter, req *http.Request) {
	dataset1 := new(data)
	to := req.URL.Query()["to"];
	from := req.URL.Query()["from"];
	amount := req.URL.Query()["amt"];
	if to == nil {
		d.Write([]byte("Error: see Help documents for request patterns"))
	}
	if to != nil {
		dataset1.to = to[0]
		dataset1.from = from[0]
		dataset1.amount = amount[0]
		sendrequest(*dataset1)
	}
}

func sendrequest(ds1 data) {
	var request string
	request += ("-a ccc " + ds1.from + " " + ds1.to + " " + ds1.amount)
}
