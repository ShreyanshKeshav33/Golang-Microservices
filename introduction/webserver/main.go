package main

import "net/http"

//Imports the net/http package, which provides support for HTTP clients and servers in Go. This package is used to create an HTTP server in your program.

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		/*
		   http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {: This line sets up a handler function for the "/hello" route.
		   When a request is made to the "/hello" endpoint, the function provided as the second argument is executed.

		   writer http.ResponseWriter: This parameter is used to write the response back to the client.
		   http.ResponseWriter interface provides methods for building and sending an HTTP response.

		   request *http.Request: This parameter represents the HTTP request received.
		   The http.Request type provides information about the incoming request, such as the URL, headers, and other details.
		*/
		writer.Write([]byte("Hello, world!"))
	})
	//writer is an instance of http.ResponseWriter, and Write is used to send data back to the client.

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err) //the errr at line 4 is aonly available to the scope of this condition i.e at line 12
	}
}
