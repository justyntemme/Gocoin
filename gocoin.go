package main

import (
	"net/http"
	"fmt"
	"net"
	"bufio"
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
	if to[0] == "source_url" {
		d.Write([]byte("github.com/justyntemme/cfetch\n"))
	}
	if to == nil {
		d.Write([]byte("Error: see Help documents for request patterns\n"))
	}
	if to != nil {
		dataset1.to = to[0]
		dataset1.from = from[0]
		dataset1.amount = amount[0]
		d.Write([]byte(sendrequest(*dataset1)))
	}
}

func sendrequest(ds1 data) string {
	var request string
	request += ("-a ccc " + ds1.amount + " " + ds1.from + " " + ds1.to)
	conn, err := net.Dial("tcp", "localhost:8888")
	fmt.Fprintf(conn, request)
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err == nil {
		fmt.Println(response)
		return(response)
	}
	return("ERROR")
}
