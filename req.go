package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
    resp, _ := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/classy")
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
