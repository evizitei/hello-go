package main

import (
  "github.com/ant0ine/go-json-rest/rest"
  "log"
  "net/http"
)

func main() {
  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)
  api.SetApp(rest.AppSimple(hello))
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func hello(w rest.ResponseWriter, r *rest.Request) {
  w.WriteJson(map[string]string{"Body": "Hello World!"})
}
