package requester

import (
  "net/http"
  "encoding/json"
  "time"
  "github.com/givey/givey-go/models"
  "io"
)

var accessToken string
var client http.Client

const (
  HOST string = "https://api-staging.givey.com"
  VERSION string = "v3"
)

func init () {
  accessToken = "gozNuicwGrfIVsBNFqKJ5ABDex7ghDuKft1MdMjf"
  client = http.Client{
    Timeout: time.Duration(5 * time.Second),
  }
}

func baseUrl () string {
  return HOST + "/" + VERSION
}

type Users struct {
  Users []models.User
}

type User struct {
  User models.User
}

func GetUsers (page string, limit string) ( error, Users ) {
  users := Users{}
  url := baseUrl() + "/users?access_token=" + accessToken + "&page=" + page
  if limit != "" {
    url = url + "&limit=" + limit
  }

  resp, err := client.Get(url)

  if err != nil {
    return err, users
  }

  decode(resp.Body, &users)

  defer resp.Body.Close()

  return nil, users
}

func GetUser (username string) ( error, User ) {
  user := User{}
  url := baseUrl() + "/users/" + username + "?access_token=" + accessToken

  resp, err := client.Get(url)

  if err != nil {
    return err, user
  }

  decode(resp.Body, &user)

  defer resp.Body.Close()

  return nil, user
}

func decode (io io.ReadCloser, target interface{}) {
  json.NewDecoder(io).Decode(target)
}
