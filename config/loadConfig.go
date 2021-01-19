package config

import (
    log "github.com/sirupsen/logrus"
    "errors"
    "os"
    "io/ioutil"
)

// LoadConfig function loads the configuration from a config file into a config.Config type
func LoadConfig(configFile string, conf *Config) error {
  log.Debug("config: Loading config from ", configFile)

  if configFile=="" {
    return errors.New("No config file specified")
  }

  jsonFile, err := os.Open(configFile)
  if err != nil{
     log.Error("Cannot open file: " +configFile)
     return err
  }
  defer jsonFile.Close()

  byteValue, err := ioutil.ReadAll(jsonFile)
  if err!=nil{
    log.Error("Cannot read file: " +configFile)
    return err
  }

  c, err := NewConfig(byteValue)
  if (err != nil){
    log.Error("Could not parse config file ["+configFile+"]: "+err.Error())
    return err
  }
  *conf = *c

  return nil
}
