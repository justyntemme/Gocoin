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
	help int
	source int
}

func main() {
	http.HandleFunc("/", serveResponse)
	http.ListenAndServe(":88", nil)
}

func serveResponse(d http.ResponseWriter, req *http.Request) {
	dataset1 := new(data)
	if req.URL.Query()["source"] != nil {
		d.Write([]byte("github.com/justyntemme/cfetch\n"))
	}
	if req.URL.Query()["help"] != nil {
		d.Write([]byte("<HELP PAGE> \n Params: to (target coin/currency) | from (coin/currency to convert) | amt (Amount of coin/currency to convert | api (Which coinfetch API to use) | help (1=show help) | source (1=Show source URL)\n"))
	}
	if req.URL.Query()["to"] != nil && req.URL.Query()["from"] != nil && req.URL.Query()["amt"] != nil && req.URL.Query()["api"] != nil {
		to := req.URL.Query()["to"];
		from := req.URL.Query()["from"];
		amount := req.URL.Query()["amt"];
		api := req.URL.Query()["api"];

		dataset1.to = to[0]
		dataset1.from = from[0]
		dataset1.amount = amount[0]
		dataset1.api = api[0]
		fmt.Printf(sendrequest(*dataset1))
		d.Write([]byte(sendrequest(*dataset1)))
	}
	if req.URL.Query()["source"] == nil && req.URL.Query()["help"] == nil && req.URL.Query()["to"] ==  nil {
			d.Write([]byte("help=1 To see help page\n"))
		}
}

func sendrequest(ds1 data) string {
	var request string
	request += ("-a " + ds1.api + " " + ds1.amount + " " + ds1.from + " " + ds1.to)
	conn, err := net.Dial("tcp", "localhost:8888")
	fmt.Fprintf(conn, request)
	response, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(response)
	if err == nil {
		fmt.Println(response)
		return(response)
	}
	fmt.Println(err.Error())
	return(err.Error())
}
