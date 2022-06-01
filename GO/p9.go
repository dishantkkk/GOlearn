package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: time.Minute}

	for i := 0; i < 5; i++ {
		resp, err := client.Get("http://webcode.me")
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(body))

		if err := resp.Body.Close(); err != nil {
			log.Fatalf("[i=%d] body close error: %s", i, err)
		}
	}
}
