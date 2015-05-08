package main

import (
    "encoding/json"
    "fmt"
    "os"
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
    dec := json.NewDecoder(os.Stdin)
    var d []data
    dec.Decode(&d)
    fmt.Printf("%+v\n", d[0])
}
