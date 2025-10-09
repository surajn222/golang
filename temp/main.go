package main

import (
	"fmt"
)

type service struct {
	string url
	statusCode int
}

func main() {
	// 3 services
	// health service
	// 4th service
	// monitor the 3 services.

	services := []string{"http://service1/health", "http://service2/health", "http://service3/health"}
	healthResults := []int{}
	
	wg := &sync.WaitGroup{}
	wg.Add(len(services))

	ch := make(chan service)

	for service := range services {
		go healthCheck(service, &wg, ch)
	}

	fmt.Println("%+v", healthResults)

	for i := range ch {
		fmt.Println("%+v, %+v", i.url, i.statusCode)
	}

	wg.Wait()
}


func healthCheck(url string, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	fmt.Println(service)

	t := time.After(
		time.Second * 10
	)

	select {
		case <- s:
			s := service{url, s}
			ch <- s
		case <- t:
			s := service{url, 500}
			ch <- s	
	}

	// hit each service
	response := http.Client(service) 

	// get the status code
	statusCode := response.statusCode

	s := service{url, statusCode}
	
	ch <- s
}
