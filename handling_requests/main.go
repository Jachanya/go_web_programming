package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*

This chapter focuses on the brain of web application, which is the handlers
Fucntion: Recieve and process request from the clients
We store session and cookies on the client side, because http is a stateless system
which means we require identity of the client.

*/

/*
net/http library has two parts (Client, Server)
*/

/*
In other to communicate between client and server securly, we use the https
protocol. The https is simply the combination of ssl and http
*/

/*Defining a struct type name handler*/
type MyHandler struct{}

/*
Creating a struct method with name ServeHTTP as it is required for it to be a handler
the ListerServer requires that we do that. Probably the function calls handle->ServeHTTP(responseWriter, request pointer)
*/

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello People!")
}

/*
We can create am decorator by chaining multiple function together.

	func protect(h http.HandlerFunc) http.HandlerFunc {
	 return func(w http.ResponseWriter, r *http.Request) {
	 . . .
	 h(w, r)
	 }
	}

How to use:

http.HandleFunc("/hello", protect(log(hello)))
*/
func main() {
	handler := MyHandler{}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        &handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
