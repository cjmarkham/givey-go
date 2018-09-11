package config

import (
  "os"
  "encoding/json"
)

type Config struct {
  Host string `json:host`
  Version string `json:version`
}
var config Config

func Load () (error, Config) {
  configFile, err := os.Open("config.json")
  defer configFile.Close()

  json.NewDecoder(configFile).Decode(&config)
  return err, config
}
