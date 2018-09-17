package main

import (
    "bytes"
    "net/http"
    "time"
)


func ClientPost(url string, text string) {
    data := []byte(text)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}

func ClientCheckConnection() bool {
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
    _, err := client.Get("http://clients3.google.com/generate_204")
    if err != nil {
        return false
    }
    return true
}