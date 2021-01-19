
package files

import (
  "testing"
)

// Test with no file specified to read
func TestNewFileReader( t *testing.T){

  err, _ := NewFileReader("")
  expectedError := "open : no such file or directory"
  if (err==nil){
    t.Error("Did not return error when no file provided expected non nil err actual nil err")
  } else if (err.Error() != expectedError){
    t.Errorf("Did not return correct error, expected [%s] actual [%s]", expectedError, err.Error())
  }
}
