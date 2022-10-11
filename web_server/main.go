package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
	response.Header().Set("Content-type", "text/html")
	fmt.Fprintf(response, "<h1>TEST HEADER<h1>")

}

func web_server() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/index", HttpFileHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func get_request() {

	resp, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func main() {

	//web_server()
	get_request()

}
