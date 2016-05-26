//  gocoin - Web socket API for interacting with the coinfetch python module 
//  Copyright (C) 2016 Justyn Temme
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, version 3 only.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU Affero General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
	api string
}

func main() {
	http.HandleFunc("/", serveResponse)
	http.ListenAndServe(":80", nil)
}

func serveResponse(d http.ResponseWriter, req *http.Request) {
	dataset1 := new(data)
	to := req.URL.Query()["to"];
	from := req.URL.Query()["from"];
	amount := req.URL.Query()["amt"];
	api := req.URL.Query()["api"];
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
		dataset1.api = api[0]
		d.Write([]byte(sendrequest(*dataset1)))
	}
}

func sendrequest(ds1 data) string {
	var request string
	request += ("-a " + ds1.api + " " + ds1.amount + " " + ds1.from + " " + ds1.to)
	conn, err := net.Dial("tcp", "localhost:8888")
	fmt.Fprintf(conn, request)
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err == nil {
		fmt.Println(response)
		return(response)
	}
	return("ERROR")
}
