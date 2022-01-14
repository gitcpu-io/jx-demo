package main

import (
  "fmt"
  "math/rand"
  "net/http"
  "time"
)

func main() {
  http.HandleFunc("/", HelloServer)
  fmt.Println("Server is Listening on 8081")
  http.ListenAndServe(":8081", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello ----by rubinus done---- Jenkins-X: %d", Random())
}

func Random() int {
  s := rand.New(rand.NewSource(time.Now().Unix()))
  return s.Intn(100)
}
