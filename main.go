package main

import (
  "fmt"
  "log"
  "net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/get" {
    http.Error(w, "404 not found", http.StatusNotFound)
    return
  }
  if r.Method != "GET" {
    http.Error(w, "only GET method supported at this endpoint", http.StatusNotFound)
    return
  }
  fmt.Fprintf(w, "get\n")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/post" {
    http.Error(w, "404 not found", http.StatusNotFound)
    return
  }
  if r.Method != "POST" {
    http.Error(w, "only POST method supported at this endpoint", http.StatusNotFound)
    return
  }
  fmt.Fprintf(w, "post\n")
}

func main() {
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)
  http.HandleFunc("/get", getHandler)
  http.HandleFunc("/post", postHandler)

  fmt.Printf("Starting server at port 8080\n")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
