package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func main() {
	url := "https://api.github.com/users/shigemk2/events"

	req, _ := http.NewRequest("GET", url, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()


	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
