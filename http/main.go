package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		// bs := make([]byte, 9999)
		// resp.Body.Read(bs)
		// fmt.Println(string(bs))
		io.Copy(os.Stdout, resp.Body)
	}
}
