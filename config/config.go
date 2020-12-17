
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
  Width string
  Mode string
  RegEx string
  PresentRatio int
}

func NewConfig(configJson *string) (*Config, error) {
  var c Config
  err := json.Unmarshal( []byte(*configJson), &c )
  if (err != nil){
    return &c, err
  } else {
      return &c, nil
  }
}
