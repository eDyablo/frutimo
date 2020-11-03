package main

import (
  "encoding/json"
  "html/template"
  "log"
  "net/http"
  "path"
)

func main() {
  http.HandleFunc("/", GetMain)
  http.HandleFunc("/info", GetInfo)
  log.Fatal(http.ListenAndServe(":80", nil))
}

func GetMain(w http.ResponseWriter, r *http.Request) {
  mainPath := path.Join("templates", "main.html")
  mainTmpl, parseErr := template.ParseFiles(mainPath)
  if parseErr != nil {
    http.Error(w, parseErr.Error(), http.StatusInternalServerError)
    return
  }
  execErr := mainTmpl.Execute(w, nil)
  if execErr != nil {
    http.Error(w, execErr.Error(), http.StatusInternalServerError)
  }
}

type Info struct {
  Name string `json:"name"`
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
  info := Info { "Frutimo" }
  jsobj, marshalErr := json.Marshal(info)
  if marshalErr != nil {
    http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(jsobj)
}
