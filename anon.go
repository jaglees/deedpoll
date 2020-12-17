package main

import (
  "flag"
  "fmt"
  "github.com/jaglees/go-anon/config"
)

func main(){
  // Capture command line switches to tool
  cmdHelp := flag.Bool("help", false, "help" )
  cmdType := flag.String("type","", "type" )
  cmdRules := flag.String("rules","", "rules" )
  cmdOutputDir := flag.String("output", "./output/", "type")
  var config *config.Config
  flag.Parse()


  if (*cmdHelp){
    displayHelp()
  } else if (len(flag.Args()) < 1 ){
    fmt.Println("You must enter the name of at least one file to process, use -help to see further details")
  } else if (*cmdType == ""){
      fmt.Println("You must define the type of file being processed, use -help to see further details")
  } else {

    var configFile string
    fmt.Println("Params: Rules=[", *cmdRules, "] Type=[", *cmdType, "] Output=[" , *cmdOutputDir, "]")
    if (*cmdRules == ""){
      configFile= fmt.Sprintf("%s.cf", *cmdType)
    } else {
      configFile= *cmdRules
    }
    config = loadConfig( configFile )

    fmt.Println("--- Configuration ---")
    fmt.Println(config)
  }
}

func loadConfig(configFile string) *config.Config {
  fmt.Println("Loading config from ", configFile)
  s := `{"Type":"customer","Header": true , "delimiter" : "," ,  "fields": [] }`
  config, err := config.NewConfig(&s)
  if (err != nil){
    panic("Could not parse config file ["+configFile+"]: "+err.Error())
  }
  return config
}

func displayHelp(){
  fmt.Println("Help!")
}
