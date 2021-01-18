package main

import (
  "flag"
  "fmt"
  "github.com/jaglees/go-anon/config"
  log "github.com/sirupsen/logrus"
)

func main(){
  // Capture command line switches to tool
  cmdHelp := flag.Bool("help", false, "Loads the help screen" )
  cmdType := flag.String("type","", "The name of the type of file (e.g. customers)" )
  cmdRules := flag.String("rules","", "(optional) The name of the rules file (defaults to the name of the type with a .cf extension e.g. customers.cf)" )
  cmdVerbose := flag.Bool("verbose", false, "(optional) Determines if debug logs are produced, outputting more details during the execution")
  cmdOutputDir := flag.String("output", "./output/", "(optional) The folder to put the resulting files")
  flag.Parse()

  // Declare variables and initialise
  var conf config.Config
  initVariables(*cmdVerbose)

  // Check to see if mandatory parameters and flags are set
  if (*cmdHelp){
    displayHelp()
  } else if (len(flag.Args()) < 1 ){
    fmt.Println("You must enter the name of at least one file to process, use -help to see further details")
  } else if (*cmdType == ""){
    fmt.Println("You must define the type of file being processed, use -help to see further details")
  } else {
    var configFile string
    log.Debug("main: Params... Rules=[", *cmdRules, "] Type=[", *cmdType, "] Output=[" , *cmdOutputDir, "]")

    // Determine the correct config file to use for this job and load the config.
    if (*cmdRules == ""){
      configFile= fmt.Sprintf("%s.cf", *cmdType)
    } else {
      configFile= *cmdRules
    }

    log.Debug("main: Config pre-loaded = [",conf,"]")
    config.LoadConfig( configFile, &conf )
    log.Debug("main: Config loaded = [",conf,"]")
  }
}

func displayHelp(){
  // Todo: need to add full help text
  fmt.Println("In order to use the anon program you must specify at least one file along with following parameters ")
  flag.PrintDefaults()
}

func initVariables(debugLogging bool){
  // log.SetFormatter(&log.TextFormatter{})
  if (debugLogging){
    log.SetLevel(log.DebugLevel)
  }
}
