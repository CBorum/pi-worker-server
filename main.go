package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	for true {
		msg := os.Getenv("msg")
		if msg == "" {
			msg = "cdcd"
		}
		resp, err := Get("http://206.189.101.6:3000/" + msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("resp:", string(resp))
		}

		time.Sleep(10 * time.Second)
	}
}

func ReadBody(resp *http.Response, err error) ([]byte, error) {
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

// Get - simple get request
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	return ReadBody(resp, err)
}
