// main.go

package main

import (
  "log"
  "net"
  "net/http"
  "fmt"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello World")
}

func main() {
  http.HandleFunc("/greeting", Greeting)

  log.Println("Starting server....")

  listener, err := net.Listen("tcp", "localhost:8080")
  if err != nil {
    log.Fatal(err)
  }

  http.Serve(listener, nil)
}
