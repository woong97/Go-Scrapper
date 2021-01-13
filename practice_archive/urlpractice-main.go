package main

import (
	"errors"
	"fmt"
	"net/http"
)

// func main() {
// 	c := make(chan string)
// 	people := [4]string{"nico", "flynn", "dal", "japanguy"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	// fmt.Println(c)
// 	for i := 0; i < len(people); i++ {
// 		fmt.Println(c)
// 		result := <-c
// 		fmt.Println(result, " ", &result)
// 	}
// 	// time.Sleep(time.Second * 10)
// }
// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 5)
// 	c <- person + "is sexy"
// }

var errRequestFailed = errors.New("Request failed")

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.google.com/",
		"https://www.facebook.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
