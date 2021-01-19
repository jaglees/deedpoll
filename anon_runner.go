package main

import (
  // "github.com/jaglees/go-anon/config"
  log "github.com/sirupsen/logrus"
  "github.com/jaglees/deedpoll/files"
)

// processFiles function processes the files from the global inputFiles slice using the logic
// which is contained in the global conf variable
func processFiles (){
  for _, f := range inputFiles {
    log.Info("main: Processing file ", f)
    processFile(f)
  }
}

func processFile(f string){
  err, fileReader := files.NewFileReader(f)
  defer fileReader.File.Close()
  if (err != nil){
    log.Warn("Cannot open file: ", f)
  }
  log.Debugf("main.processFiles: opened file %s starting reading", f)

  for fileReader.Scanner.Scan() {
    log.Info(fileReader.Scanner.Text())
  }
}
