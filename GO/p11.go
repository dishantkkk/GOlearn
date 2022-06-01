package main

import (
    "log"
    "net/http"
    "os"
    "io"
    "fmt"
)

func main() {

    r, err := http.Get("http://webcode.me")

    if err != nil {
      log.Fatal(err)
    }

    defer r.Body.Close()

    f, err := os.Create("index-f-go.html")

	f, err = os.Create("index-f-go1.html")

	h, err := os.Create("index-f-go2.html")

    if err != nil {
      log.Fatal(err)
    }

    defer f.Close()

	defer h.Close()

    _, err = f.ReadFrom(r.Body)
    body, err := io.ReadAll(r.Body)
    fmt.Println(string(body))
	_, err = h.ReadFrom(r.Body)

    if err != nil {
      log.Fatal(err)
    }
}
