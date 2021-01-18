package files

import (
  log "github.com/sirupsen/logrus"
  "sync"
)


func ReadLines(fileName string, c chan string, wg *sync.WaitGroup) error {
  log.Debug("files ReadLines: Reading lines and writing to channel")

  // Read all lines from file and write to outChan
  c <- "Hello"

  close(c)
  log.Debug("files ReadLines: Written")
  wg.Done()
  return nil
}
