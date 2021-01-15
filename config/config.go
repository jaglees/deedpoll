
package config

import (
  "encoding/json"
)

type Config struct {
  Type string
  Header bool
  Delimiter string
  Fields []Field
}

type Field struct {
  Name string
  Width int
  Mode string
  RegEx string
  PresentRatio float32
}

func NewConfig(configJson []byte) (*Config, error) {
  var c Config
  err := json.Unmarshal( configJson, &c )
  if (err != nil){
    return &c, err
  } else {
      return &c, nil
  }
}
