package main

import (
	"encoding/json"
	"fmt"
)

type PoliticsJson struct {
}

// type PoliticsJson struct {
// 	Results []struct {
// 		Multimedia []struct {
// 			URL string `json:"url"`
// 		} `json:"multimedia"`
// 		Title string `json:"title"`
// 	} `json:"results"`
// }

func main() {
	s := ``

	var p PoliticsJson

	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%T", p.Results)
	// fmt.Println(p.Results[0].Title)
	// fmt.Println(p.Results[0].Multimedia[0].URL)

	var strSlice []string

	for i := 0; i < len(p); i++ {
		// fmt.Println(p.Results[i].Title)
		// fmt.Printf("%T\n", p.ThreatClass[i].ID)
		strSlice = append(strSlice, p[i].ID)
	}

	fmt.Println(strSlice)
}
