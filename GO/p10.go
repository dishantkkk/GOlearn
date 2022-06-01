package main
import (
    "log"
    "io/ioutil"
    "net/http"
    "os"
)
func main() {
    r, err := http.Get("http://webcode.me")
    if err != nil {
      log.Fatal(err)
    }

    defer r.Body.Close()
    f, err := os.Create("index.html")
    if err != nil {
      log.Fatal(err)
    }
    ioutil.WriteFile("index1.html", body)
    defer f.Close()
    _, err = f.ReadFrom(r.Body)

    if err != nil {
      log.Fatal(err)
    }
}
