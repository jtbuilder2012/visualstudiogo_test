/*
This programm demonstrates how to build a simple http server using go

*/

package main

import (
	"fmt"
	"net/http"
)

func main() {

	//Define server handler for directories configured on the http server
	//To test open the browser at localhost:8085/ and localhost:8085/earth
	//
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)

	//Define the server port of whic to listen for requests
	//To interact with this server open the browser with address: localhost:8080
	//Manage any local port conflicts, e.g the port is assigned to another server
	http.ListenAndServe(":8085", nil)

}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World - http server")
}
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Earth World - http server")
}
