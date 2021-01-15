package config

import (
    log "github.com/sirupsen/logrus"
    "errors"
    "os"
    "io/ioutil"
)

func LoadConfig(configFile string, conf *Config) error {
  log.Debug("Loading config from ", configFile)

  if configFile==""{
    return errors.New("No config file specified")
  }

  jsonFile, err := os.Open(configFile)
  if err != nil{
     log.Error("Cannot open file: " +configFile)
     return err
  }

  byteValue, err := ioutil.ReadAll(jsonFile)
  if err!=nil{
    log.Error("Cannot read file: " +configFile)
    return err
  }

  // s := `{"Type":"customer","Header": true , "delimiter" : "," ,  "fields": [        {
  //             "name": "lineNo",
  //             "mode": "origional",
  //             "presentRatio": 1
  //         },
  //         {
  //             "name": "customerId",
  //             "mode": "token",
  //             "regEx": "[A-Z]{3}[0-9]{3}",
  //             "presentRatio": 1
  //         }] }`

  c, err := NewConfig(byteValue)
  if (err != nil){
    log.Error("Could not parse config file ["+configFile+"]: "+err.Error())
    return err
  }
  *conf = *c

  return nil
}
