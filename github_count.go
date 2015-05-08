package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type actor struct {
    Id   int
    Login  string
    GravatarId string
    Url string
    AvatarUrl string
}

type repo struct {
    Id   int
    Name string
    Url string `json:"url"`
}

type payload struct {
    Ref   string
    RefType  string
    MasterBranch string
    Description string
    PusherType string
}

type data struct {
    Id      string `json:"id"`
    Type      string `json:"type"`
    Actor      actor `json:"actor"`
    Repo      repo `json:"repo"`
    PayLoad      payload `json:"payload"`
    Public       string `json:"public"`
    CreatedAt      string `json:"created_at"`
}

func main() {
	url := "https://api.github.com/users/shigemk2/events"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var d []data
	dec.Decode(&d)
	fmt.Printf("%+v\n", d[0].CreatedAt)
}
