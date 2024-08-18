package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func SimpleTextResponse(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "SimpleTextResponse")
}

func SimpleListResponse(response http.ResponseWriter, request *http.Request) {
	list1 := []string{"element1", "element2"}
	fmt.Fprintf(response, strings.Join(list1, ","))
}

func SimpleMapResponse(response http.ResponseWriter, request *http.Request) {
	map_2 := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
	}
	mJson, _ := json.Marshal(map_2)
	fmt.Fprintf(response, string(mJson))
}

func RequestComponents(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "<h1>TEST HEADER<h1>")
}

func JsonResponse(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "<h1>TEST HEADER<h1>")
}

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
	response.Header().Set("Content-type", "text/html")
	fmt.Fprintf(response, "<h1>TEST HEADER<h1>")

}

func web_server() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hi")
	// })

	http.HandleFunc("/simpleTextResponse", SimpleTextResponse)
	http.HandleFunc("/simpleListResponse", SimpleListResponse)
	http.HandleFunc("/simpleMapResponse", SimpleMapResponse)
	http.HandleFunc("/simpleFileResponse", HttpFileHandler)
	http.HandleFunc("/requestComponents", RequestComponents)
	http.HandleFunc("/jsonResponse", JsonResponse)

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
	web_server()
	//get_request()

}
