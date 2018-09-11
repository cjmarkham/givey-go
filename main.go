package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "github.com/givey/givey-go/requester"
)

func main () {
  requester.Init()
  router := mux.NewRouter()
  router.HandleFunc("/users/{username:[a-zA-Z0-9]+}", getUser)
  router.HandleFunc("/users", getUsers)
  http.Handle("/", router)
  http.ListenAndServe(":3005", nil)
}

func getPage (request *http.Request) string {
  url := request.URL
  query := url.Query()

  page := "1"
  if query.Get("page") != "" {
    page = query.Get("page")
  }

  return string(page)
}

func getLimit (request *http.Request) string {
  url := request.URL
  query := url.Query()

  limit := query.Get("limit")

  return string(limit)
}

func getUsers (writer http.ResponseWriter, request *http.Request) {
  page := getPage(request)
  limit := getLimit(request)

  err, userData := requester.GetUsers(page, limit)
  if err != nil {
    http.Error(writer, err.Error(), http.StatusInternalServerError)
    return
  }
  writer.Header().Set("Content-Type", "application/json")
  js, err := json.Marshal(userData)
  if err != nil {
    http.Error(writer, err.Error(), http.StatusInternalServerError)
    return
  }
  writer.Write(js)
}

func getUser (writer http.ResponseWriter, request *http.Request) {
  vars := mux.Vars(request)
  username := vars["username"]

  err, userData := requester.GetUser(username)
  if err != nil {
    http.Error(writer, err.Error(), http.StatusInternalServerError)
    return
  }
  writer.Header().Set("Content-Type", "application/json")
  js, err := json.Marshal(userData)
  if err != nil {
    http.Error(writer, err.Error(), http.StatusInternalServerError)
    return
  }
  writer.Write(js)
}
