package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	// "encoding/json"
)


func main() {
	url := "https://api.github.com/users/shigemk2/events"

	req, _ := http.NewRequest("GET", url, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()


	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(byteArray)
	// fmt.Println(string(byteArray))
}
