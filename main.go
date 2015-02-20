package main

import (
  "fmt"
  "log"
  "net/http"
  "io/ioutil"
  "strings"
)

func main() {
  fmt.Println("Begin main!\n\n")

  resp, err := http.Get("http://en.wikipedia.org/wiki/Lost_(TV_series)")
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  sBody := strings.Split(string(body), " ")

  words := make([]int, 30, 30)
  tags := 0
  for _,element := range sBody {
    if len(element) > 0 && len(element) < 30 { //&& strings.HasPrefix(element, "<")
      words[len(element)] = words[len(element)] + 1
    }

    if strings.HasPrefix(element, "<") {
      tags++
    }

  }

  fmt.Println("Words size:")
  fmt.Println(words)

  fmt.Printf("\nHtml tags: %d\n", tags)

  //fmt.Println(string(body))
  fmt.Println("\n\nEnd main!")
}
