package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://ipinfo.io")
	if err != nil {
		fmt.Printf("Request failed with error %s\n", err.Error())
	} else {
		r, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%d - %s\n", resp.StatusCode, r)
	}
}
