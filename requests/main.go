package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "golang.org/x/exp/slices"
)

func main() {

	// list_to_check := []int{234, 567}
	list_threat_class, _ := get_list()
	fmt.Println(list_threat_class)

	// for _, v := range list_to_check {
	// 	slices.Contains(list_threatClass, v)
	// 	fmt.Println(v)
	// }

}

func get_list() ([]int, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	url := "http://httpbin.org/ip"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	// req.Header.Add("Authorization", "Bearer TOKEN")

	resp, err := client.Do(req)

	// Throw error when the response is not right
	if err != nil {
		log.Fatalln(err)
	}
	// Convert the response to list first
	//We Read the response body on the line below.
	type PoliticsJson struct {
		Origin string `json:"origin"`
		Str1   string `json:"str1"`
		Str2   string `json:"str2"`
	}

	var p PoliticsJson

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal([]byte(body), &p)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%T", p.Results)
	// fmt.Println(p.Results[0].Title)
	// fmt.Println(p.Results[0].Multimedia[0].URL)

	var strSlice []string

	for i := 0; i < len(p.Origin); i++ {
		// fmt.Println(p.Results[i].Title)
		// fmt.Printf("%T\n", p.ThreatClass[i].ID)
		strSlice = append(strSlice, p.Origin)
	}

	// for i := 0; i < len(p.ThreatClass); i++ {
	// 	// fmt.Println(p.Results[i].Title)
	// 	// fmt.Printf("%T\n", p.ThreatClass[i].ID)
	// 	strSlice = append(strSlice, p.ThreatClass[i].ID)
	// }

	fmt.Println(strSlice)

	// Throw error when unable to convert to list
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	s1 := []int{234, 567, 7890, 1234, 234}

	return s1, nil
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero", a)
	}
	return a / b, nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// go get golang.org/x/exp/slices
// import  "golang.org/x/exp/slices"
// things := []string{"foo", "bar", "baz"}
// slices.Contains(things, "foo") // true

// os.Setenv("FOO", "1")
// fmt.Println("FOO:", os.Getenv("FOO"))
// fmt.Println("BAR:", os.Getenv("BAR"))
