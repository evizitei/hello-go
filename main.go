package main

import (
  "github.com/ant0ine/go-json-rest/rest"
  "log"
  "net/http"
)

func main() {
  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)

  router, err := rest.MakeRouter(
    rest.Get("/", hello),
    rest.Post("/sum", sum),
  )

  if err != nil {
    log.Fatal(err)
  }

  api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func hello(w rest.ResponseWriter, r *rest.Request) {
  w.WriteJson(map[string]string{"Body": "Hello World! Post a JSON array of numbers (key = \"Numbers\") to /sum for summing action"})
}

type SumInput struct {
  Numbers []int
}

type SumOutput struct {
  Solution int
}

func sum(w rest.ResponseWriter, r *rest.Request) {
    input := SumInput{}
    err := r.DecodeJsonPayload(&input)
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if len(input.Numbers) <= 0 {
        rest.Error(w, "numbers needed for summing", 400)
        return
    }
    sum := 0
    for val := range input.Numbers {
        sum += val
    }
    output := SumOutput{
      Solution: sum,
    }
    w.WriteJson(&output)
}
