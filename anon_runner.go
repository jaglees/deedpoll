package main

import (
  // "github.com/jaglees/go-anon/config"
  log "github.com/sirupsen/logrus"
  "github.com/jaglees/go-anon/files"
  "sync"
)

// processFiles function processes the files from the global inputFiles slice using the logic
// which is contained in the global conf variable
func processFiles (){
  inputChan := make(chan string)
  // outputChan := make(chan string)

  for _, f := range inputFiles {
    log.Info("Processing file ", f)
    var wg sync.WaitGroup
  	wg.Add(2)

    go files.ReadLines("", inputChan, &wg)
    go files.WriteLines("", inputChan, &wg)
    wg.Wait()
    log.Debug("Waitgroup complete")
    // file reader
    // anonocd miser
    // file writer
  }
}
