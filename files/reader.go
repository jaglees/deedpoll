package files

import (
  log "github.com/sirupsen/logrus"
  "bufio"
  "os"
)

type FileReader struct {
  File *os.File
  Scanner *bufio.Scanner
}

func NewFileReader(filename string) (error, FileReader) {
  log.Debug("files NewFileReader: creating a new FileReader type")

  var fr FileReader
  var err error
  fr.File, err = os.Open(filename)
  if err != nil {
    return err, fr
  }

  // Start reading from the file with a reader.
  fr.Scanner = bufio.NewScanner(fr.File)

  log.Debug("files OpenFile: finished")
  return nil, fr
}
