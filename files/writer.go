package files

import (
  log "github.com/sirupsen/logrus"
  "sync"
)

func WriteLines(fileName string, c chan string, wg *sync.WaitGroup) error {
  log.Debug("files WriteLines: Writing lines from channel to file")
  for i := range c {
		log.Debug(i)
	}
  log.Debug("files WriteLines: Read")
  wg.Done()
  return nil
}
