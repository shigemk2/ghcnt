package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func main() {
	url := "https://api.github.com/users/shigemk2/events"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
