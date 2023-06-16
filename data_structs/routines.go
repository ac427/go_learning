package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com/",
		"https://failme.foo.ac",
	}

	c := make((chan string))

	for _, url := range urls {
		go checkStatus(url, c)
	}

	// blocking channel
	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkStatus(<-c, c)
	// }

	// same as above code for {}
	// for l := range c {
	// 	// time.Sleep(5* time.Second) . The main routine is going to get blocked and may not work as we expect
	// 	go checkStatus(l, c)

	// }

	for l := range c {
		// function literal. similar to lamda in python or anonymous func in js
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkStatus(link, c)
		}(l)

	}

}

func checkStatus(url string, c chan string) {
	resp, _ := http.Get(url)
	if resp != nil {
		fmt.Println(url, " status is: ", resp.StatusCode)
		// c <- "Got some response"
		c <- url
	} else {
		fmt.Println(url, " Timeout/Failed to Get")
		// c <- "failed, no response"
		c <- url
	}
}
