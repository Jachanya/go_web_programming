package main

/*Http messages are two type Request and Response*/
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sau Sheong",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

// Write()
// WriteHeader
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	url := r.URL
	h_res := w.Header
	fmt.Fprintln(w, *url)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
	fmt.Fprintln(w, h_res)
	fmt.Fprintln(w, h)
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello World</h1></body>
	</html>`
	w.Write([]byte(str))
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/headers", headers)
	http.HandleFunc("/json", jsonExample)
	http.HandleFunc("/set_cookie", setCookie)

	server.ListenAndServe()
}

/*Cookies*/

/*cookies are a way to simulate persistence*/
