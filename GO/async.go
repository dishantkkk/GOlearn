package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "regexp"
  "sync"
)

func main() {

  urls := []string{
    "http://webcode.me",
    "https://example.com",
    "http://httpbin.org",
    "https://www.perl.org",
    "https://www.php.net",
    "https://www.python.org",
    "https://code.visualstudio.com",
    "https://clojure.org",
  }

  var wg sync.WaitGroup

  for _, u := range urls {

    wg.Add(1)
    go func(url string) {

      defer wg.Done()

      content := doReq(url)
      title := getTitle(content)
      fmt.Println(title)
    }(u)
  }

  wg.Wait()
}



  func doReq (url string) (content string) {

  resp, err := http.Get(url)
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatal(err)
  }

  return string(body)
  }

  func getTitle(content string) (title string) {

  re := regexp.MustCompile("<title>(.*)</title>")

  parts := re.FindStringSubmatch(content)

  if len(parts) > 0 {
    return parts[1]
  } else {
    return "no title"
  }
}
