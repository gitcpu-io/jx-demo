package main

import (
  "fmt"
  "math/rand"
  "net/http"
  "os"
  "time"
)

var commitOld = "latest"
var commit = "latest"

func init() {
  if os.Getenv("COMMIT") != "" {
    commit = os.Getenv("COMMIT")
  }
  if os.Getenv("COMMIT_OLD") != "" {
    commitOld = os.Getenv("COMMIT_OLD")
  }
}

func main() {
  http.HandleFunc("/", HelloServer)
  fmt.Println("Server is Listening on 8081")
  http.ListenAndServe(":8081", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Old's commit id: %s </h1><h1>PR's commit id: %s </h1><h1>Hello -- Jenkins-X: %d</h1>", commitOld, commit, Random())
}

func Random() int {
  s := rand.New(rand.NewSource(time.Now().Unix()))
  return s.Intn(100)
}
