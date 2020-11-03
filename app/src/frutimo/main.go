package main

import (
  "log"
  "net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("frutimo"))
}

func main() {
  http.HandleFunc("/", GetIndex)
  log.Fatal(http.ListenAndServe(":80", nil))
}
